package dev

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"io"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	s "strings"
	"time"
)

type Manager struct {
	Client   *client.Client
	DockSock string
}
type DockerComposeConfig struct {
	Version string `yaml:"version"`

	Services map[string]ContainerBuild `yaml:"services,omitempty"`
	//Volumes  map[string]Volume  `yaml:"volumes,omitempty"`
	//Networks map[string]Network `yaml:"networks,omitempty"`
	//Secrets  map[string]Secret  `yaml:"secrets,omitempty"`
	//Configs  map[string]Config  `yaml:"configs,omitempty"`
}

type ContainerBuild struct {
	ContainerName string                   `yaml:"container_name"`
	ContainerId   string                   `yaml:"container_id"`
	Config        container.Config         `yaml:"config"`
	HostConfig    container.HostConfig     `yaml:"host_config"`
	NetworkConfig network.NetworkingConfig `yaml:"network_config"`
	Platform      v1.Platform              `yaml:"platform"`
}

func CheckColima() error {
	err := exec.Command("colima", "status").Run()
	if err != nil {
		return fmt.Errorf("error checking Colima status -> %s", err.Error())
	}
	return nil
}

func GetUser() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting current user -> %s", err.Error())
	}
	return usr.Username, nil
}

func ValidateBuild(build ContainerBuild) error {
	errString := "build validation error:"

	if build.Config.Image == "" {
		return fmt.Errorf("%s image name dne -> %s", errString, build.Config.Image)
	}

	if build.ContainerName == "" {
		return fmt.Errorf("%s container name dne", errString)
	}
	return nil
}

func (d *Manager) RunContainer(ctx context.Context, build ContainerBuild) (string, error) {
	// Creating the actual container. This is "nil,nil,nil" in every example.
	cont, err := d.Client.ContainerCreate(
		ctx,
		&build.Config,
		&build.HostConfig,
		&build.NetworkConfig,
		&build.Platform,
		build.ContainerName,
	)
	if err != nil {
		return "", err
	}

	// Run the actual container
	err = d.Client.ContainerStart(context.Background(), cont.ID, container.StartOptions{})
	if err != nil {
		return "", err
	}

	return cont.ID, nil
}

func (d *Manager) ListContainers(ctx context.Context) (containers []types.Container, err error) {
	containers, err = d.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return containers, fmt.Errorf("error listing containers -> %s", err.Error())
	}
	return
}

func (d *Manager) GetContainer(ctx context.Context, nameOrId string) (ct types.Container, err error) {
	var containers []types.Container
	containers, err = d.ListContainers(ctx)
	if err != nil {
		return ct, fmt.Errorf("error getting container -> %s", err.Error())
	}

	for _, c := range containers {
		if c.ID == nameOrId {
			return c, nil
		}
		for _, name := range c.Names {
			if s.Contains(name, nameOrId) {
				return c, nil
			}
		}
	}
	return ct, errors.New("container doesn't exist")
}

func (d *Manager) ClientConnect() error {
	var err error

	if runtime.GOOS == "darwin" {
		uID, err := GetUser()
		if err != nil {
			return fmt.Errorf("failed to get user: %v", err)
		}
		d.DockSock = fmt.Sprintf("unix:///Users/%s/.colima/default/docker.sock", uID)
	} else if runtime.GOOS == "linux" {
		d.DockSock = "unix:///var/run/docker.sock"
	}

	err = os.Setenv("DOCKER_HOST", d.DockSock)
	if err != nil {
		return fmt.Errorf("failed to set DOCKER_HOST: %v", err)
	}

	d.Client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}

	for i := 0; i < 3; i++ {
		_, err := d.Client.Ping(context.TODO())
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	return nil
}

func (d *Manager) Pull(ctx context.Context, image string) error {
	reader, err := d.Client.ImagePull(ctx, image, types.ImagePullOptions{})

	defer reader.Close()
	if err != nil {

		println(err.Error())
		return err
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func (d *Manager) DeleteContainer(ctx context.Context, id string) error {
	if err := d.Client.ContainerStop(ctx, id, container.StopOptions{}); err != nil {
		return err
	}
	err := d.Client.ContainerRemove(ctx, id, types.ContainerRemoveOptions{RemoveVolumes: true, Force: true})
	if err != nil {
		return err
	}
	return nil
}

func (d *Manager) FindCachedImage(ctx context.Context, imgName string) (types.ImageInspect, error) {
	images, err := d.Client.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return types.ImageInspect{}, err
	}
	for _, i := range images {
		for _, tag := range i.RepoTags {
			if tag == imgName {
				inspect, _, err := d.Client.ImageInspectWithRaw(ctx, i.ID)
				if err != nil {
					return types.ImageInspect{}, err
				}
				return inspect, nil
			}
		}
	}
	return types.ImageInspect{}, nil
}

func (d *Manager) FreeUsedPort(ctx context.Context, port string) error {
	iPort, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port number: %s", port)
	}
	cList, err := d.Client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return err
	}

	for _, containers := range cList {
		for _, portInfo := range containers.Ports {
			if portInfo.PublicPort == uint16(iPort) {
				fmt.Printf("Found container %s using port %s, removing container\n", containers.Names, port)
				err = d.DeleteContainer(ctx, containers.ID)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	return fmt.Errorf("no running container found on port %s", port)
}

func (d *Manager) PruneContainer(ctx context.Context) error {
	fil := filters.Args{}
	_, err := d.Client.ContainersPrune(ctx, fil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Manager) BuildImage(ctx context.Context, tags []string, dockerfile string) error {
	// Create a buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Create a filereader
	dockerFileReader, err := os.Open(dockerfile)
	if err != nil {
		return err
	}

	// Read the actual Dockerfile
	readDockerFile, err := io.ReadAll(dockerFileReader)
	if err != nil {
		return err
	}

	// Make a TAR header for the file
	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readDockerFile)),
	}

	// Writes the header described for the TAR file
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}

	// Writes the dockerfile data to the TAR file
	_, err = tw.Write(readDockerFile)
	if err != nil {
		return err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	// Build the actual image
	imageBuildResponse, err := d.Client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

func (d *Manager) GetLatestCachedImgVersion(ctx context.Context, imgName string) (types.ImageInspect, error) {
	images, err := d.Client.ImageList(ctx, types.ImageListOptions{All: true})
	var latestTag string
	var latestImg types.ImageInspect
	if err != nil {
		return latestImg, err
	}

	for _, i := range images {
		for _, tag := range i.RepoTags {
			if s.Contains(tag, imgName) && latestTag < tag {
				fmt.Printf("%v\n", tag)
				latestTag = tag

				inspect, _, err := d.Client.ImageInspectWithRaw(ctx, i.ID)
				if err != nil {
					return types.ImageInspect{}, err
				}
				latestImg = inspect
			}
		}
	}

	return latestImg, nil
}

func (d *Manager) Exec(ctx context.Context, containerName string, cmd []string) {
	// Create a Docker client with specified API version
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Error creating Docker client:", err)
		return
	}

	containerID, err := d.GetContainerID(ctx, cli, containerName)
	if err != nil {
		fmt.Println("Error getting container ID:", err)
		return
	}

	execCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	})
	if err != nil {
		fmt.Println("Error creating exec instance:", err)
		return
	}

	resp, err := cli.ContainerExecAttach(ctx, execCreateResp.ID, types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	})
	if err != nil {
		fmt.Println("Error attaching to exec instance:", err)
		return
	}
	defer resp.Close()

	// Print the command output
	output, err := io.ReadAll(resp.Reader)
	if err != nil {
		fmt.Println("Error reading exec instance output:", err)
		return
	}

	fmt.Printf("Output:\n%s\n", output)

	execInspectResp, err := cli.ContainerExecInspect(ctx, execCreateResp.ID)
	if err != nil {
		fmt.Println("Error inspecting exec instance:", err)
		return
	}

	fmt.Printf("\nExit Code: %d\n", execInspectResp.ExitCode)

}

func (d *Manager) GetContainerID(ctx context.Context, cli *client.Client, containerName string) (string, error) {
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", err
	}

	for _, tainer := range containers {
		for _, name := range tainer.Names {
			if s.Contains(name, containerName) {
				return tainer.ID, nil
			}
		}
	}

	return "", fmt.Errorf("container not found: %s", containerName)
}

func NewContainerBuilder() ContainerBuild {

	return ContainerBuild{
		HostConfig: container.HostConfig{
			PortBindings: nat.PortMap{},
			LogConfig: container.LogConfig{
				Type:   "json-file",
				Config: map[string]string{},
			},
			//RestartPolicy: container.RestartPolicy{ //todo make configurable
			//	Name: "always",
			//},
		},
		Config: container.Config{
			ExposedPorts: nat.PortSet{},
			AttachStderr: true,
			AttachStdout: true,
			//User:         "4000:4000", //todo make configurable
		},
		//NetworkConfig: network.NetworkingConfig{
		//	EndpointsConfig: map[string]*network.EndpointSettings{
		//		"bridge": {
		//			Gateway: "gatewayname",
		//		},
		//	},
		//},
	}
}

/*
// todo figure out root settings
// gets tricky see Postgres issue with setting root
// https://www.reddit.com/r/docker/comments/v89opx/run_postgres_container_as_non_root_user_volume/
//This optional variable can be used to define another location - like a subdirectory -
for the database files. The default is /var/lib/postgresql/data.
If the data volume you're using is a filesystem mountpoint (like with GCE persistent disks),
or remote folder that cannot be chowned to the postgres user (like some NFS mounts),
or contains folders/files (e.g. lost+found), Postgres initdb requires a subdirectory to
be created within the mountpoint to contain the data.
*/

// todo finish build template
func NewNonRootContainerBuild() ContainerBuild {
	return ContainerBuild{
		HostConfig: container.HostConfig{
			PortBindings: nat.PortMap{},
			LogConfig: container.LogConfig{
				Type:   "json-file",
				Config: map[string]string{},
			},
			//RestartPolicy: container.RestartPolicy{ //todo make configurable
			//	Name: "always",
			//},
		},
		Config: container.Config{
			ExposedPorts: nat.PortSet{},
			AttachStderr: true,
			AttachStdout: true,
			//User:         "4000:4000", //todo make configurable
		},
		//NetworkConfig: network.NetworkingConfig{
		//	EndpointsConfig: map[string]*network.EndpointSettings{
		//		"bridge": {
		//			Gateway: "gatewayname",
		//		},
		//	},
		//},
	}
}

func (d *Manager) CreateNetwork(ctx context.Context, networkName string, networkBuild types.NetworkCreate) (nr types.NetworkCreateResponse, err error) {
	nr, err = d.Client.NetworkCreate(ctx, networkName, networkBuild)
	if err != nil {
		return nr, fmt.Errorf("failed to create network -> %s", err.Error())
	}
	return
}

func (d *Manager) CleanCreateNetwork(ctx context.Context, networkName string, networkBuild types.NetworkCreate) (types.NetworkCreateResponse, error) {
	nr := types.NetworkCreateResponse{}
	err := d.DeleteNetwork(ctx, networkName)
	if err != nil {
		return types.NetworkCreateResponse{}, err
	}

	nr, err = d.CreateNetwork(ctx, networkName, networkBuild)
	if err != nil {
		return nr, fmt.Errorf("failed to create newtwork -> %s", err.Error())
	}
	return nr, nil
}

func (d *Manager) DeleteNetwork(ctx context.Context, networkName string) error {
	ni, err := d.Client.NetworkInspect(ctx, networkName, types.NetworkInspectOptions{})
	if err != nil {
		if s.Contains(err.Error(), "not found") {
			return nil
		}
		return fmt.Errorf("failed to inspect network -> %s", err.Error())
	} else {
		fmt.Printf("network exists: %s, removing\n", networkName)

		for _, endpoint := range ni.Containers {
			fmt.Printf("container exists on network: %s, removing\n", endpoint.Name)
			getContainer, err := d.GetContainer(ctx, endpoint.Name)
			if err != nil {
				return fmt.Errorf("failed to get container %s -> %s", endpoint.Name, err.Error())
			}
			err = d.DeleteContainer(ctx, getContainer.ID)
			if err != nil {
				return fmt.Errorf("failed to stop/delete container %s -> %s", endpoint.Name, err.Error())
			}
		}

		err = d.Client.NetworkRemove(ctx, ni.ID)
		if err != nil {
			return fmt.Errorf("failed to remove network -> %s", err.Error())
		}
	}
	return nil
}

func (d *Manager) GracefulShutdown(ctx context.Context, networkName string) error {
	err := d.DeleteNetwork(ctx, networkName)
	if err != nil {
		return err
	}
	return nil
}

//todos
//todo make network cleanup defer func
//todo make a git.target.com module so it can be imported
//todo fix module version negotiation so module docker isn't incompatible
//todo Make NotRootContainer Template

//docker run \
//   -p 9000:9000 \
//   -p 9001:9001 \
//   --user $(id -u):$(id -g) \
//   -v ${HOME}/minio/data:/data \
//   quay.io/minio/minio server /data --console-address ":9001"

//todo Volume manager
//todo Clean env and clear docker - remove all existing containers
//todo clear images - remove all existing images
//todo Run containers -> Maybe different function with clearer name since it also frees ports and deletes older containers

//Todo make more user friendly
//	template files? how does a user configure the defaults?

//todo update running containers?
//todo  - better patter to deploy new one with changes?
//todo  + legacy containers that are not reconfigurable

//todo cancel with context
//todo compose files
//todo similar kubernetes manager app
//todo config module
//ports tracker?
//container security checker
//test containers?
//docker to kubernetes translator
//container console stream
//templates for postgres, redis, minio

//todo docker compose
//todo sha check
