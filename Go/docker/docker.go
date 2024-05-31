package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	s "strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	//"github.com/compose-spec/compose-go/cli"
	//"github.com/docker/cli/cli/command"
	//"github.com/docker/cli/cli/flags"
	//"github.com/docker/compose/v2/cmd/formatter"
	//"github.com/docker/compose/v2/pkg/api"
	//"github.com/docker/compose/v2/pkg/compose"
)

//todo docker compose
//todo sha check

type Manager struct {
	Client   *client.Client
	DockSock string
}

type ContainerBuild struct {
	ImgName       string
	ContainerName string
	Port          string
	EnvVars       []string
	Cmd           []string
	ContainerId   string
	Version       string
	CPath         string
	Volumes       []string
}

func CheckColima() error {
	err := exec.Command("colima", "status").Run()
	if err != nil {
		return err
		//println(erp.Cause(err).Error())
		//cmd := exec.Command("colima", "start")
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//err := cmd.Run()
		//if err != nil {
		//	return erp.Wrap(err, "Failed to Start Colima\n")
		//}
	}
	return nil
}

func GetUser() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}

func ValidateBuild(build ContainerBuild) error {
	errString := "build validation error:"

	if build.ImgName == "" {
		return fmt.Errorf("%s image name dne -> %s", errString, build.ImgName)
	}

	if build.Version == "" {
		return fmt.Errorf("%s image version dne -> %s", errString, build.Version)
	}

	if build.ContainerName == "" {
		return fmt.Errorf("%s container name dne", errString)
	}
	return nil
}

func (d *Manager) RunContainer(ctx context.Context, build ContainerBuild) (string, error) {
	// Define a PORT opening
	newport, err := nat.NewPort("tcp", build.Port)
	if err != nil {
		fmt.Println("unable to create nat.NewPort for container")
		return "", err
	}

	// Configured hostConfig:
	// https://godoc.org/github.com/docker/docker/api/types/container#HostConfig

	hostConfig := &container.HostConfig{
		Binds: build.Volumes,
		PortBindings: nat.PortMap{
			newport: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: build.Port,
				},
			},
		},

		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{},
		},
	}

	// Define Network config (why isn't PORT in here...?:
	// https://godoc.org/github.com/docker/docker/api/types/network#NetworkingConfig
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	// Define ports to be exposed (has to be same as hostconfig.portbindings.newport)
	exposedPorts := map[nat.Port]struct{}{
		newport: struct{}{},
	}

	// Configuration
	// https://godoc.org/github.com/docker/docker/api/types/container#Config
	config := &container.Config{
		Image:        fmt.Sprintf("%s:%s", build.ImgName, build.Version),
		Env:          build.EnvVars,
		ExposedPorts: exposedPorts,
		Hostname:     fmt.Sprintf("%s-hostnameexample", build.ImgName),
	}

	platform := &v1.Platform{}

	//d.Client.ContainerExecCreate()

	// Creating the actual container. This is "nil,nil,nil" in every example.
	cont, err := d.Client.ContainerCreate(
		ctx,
		config,
		hostConfig,
		networkConfig,
		platform,
		build.ContainerName,
	)
	if err != nil {
		return "", err
	}

	// Run the actual container
	err = d.Client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return cont.ID, nil
}

func (d *Manager) ListContainers(ctx context.Context) (containers []types.Container, err error) {
	containers, err = d.Client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return
	}
	//fmt.Printf("%-12s\t%-30s\t%-10s\t%s\n", "CONTAINER ID", "IMAGE", "STATUS", "PORTS")
	//for _, c := range containers {
	//	fmt.Printf("%-12s\t%.30s\t%.10s\t%v\n", c.ID[:12], fmt.Sprintf("%-30s", c.Image), c.Status, c.Ports)
	//}

	return
}

func (d *Manager) GetContainer(ctx context.Context, nameOrId string) (ct types.Container, err error) {
	var containers []types.Container
	containers, err = d.ListContainers(ctx)
	if err != nil {
		return
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

func (d *Manager) Connect(ctx context.Context) (err error) {
	if runtime.GOOS == "darwin" {
		uID, err := GetUser()
		if err != nil {
			return err
		}
		d.DockSock = fmt.Sprintf("unix:///Users/%s/.colima/default/docker.sock", uID)
		err = os.Setenv("DOCKER_HOST", d.DockSock)
		if err != nil {
			return err
		}
	} else if runtime.GOOS == "linux" {
		d.DockSock = fmt.Sprintf("unix:///var/run/docker.sock")
		err = os.Setenv("DOCKER_HOST", d.DockSock)
		d.Client, err = client.NewClientWithOpts()
	}

	d.Client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
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
	//TODO fix Pull Panic when no connection

	reader, err := d.Client.ImagePull(ctx, image, types.ImagePullOptions{})

	defer reader.Close()
	if err != nil {
		println(err.Error())
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		println(err.Error())
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

func (d *Manager) FindCachedImage(ctx context.Context, imgName string) (types.ImageSummary, error) {
	images, err := d.Client.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return types.ImageSummary{}, err
	}
	for _, i := range images {
		for _, tag := range i.RepoTags {
			if tag == imgName {
				return i, nil
			}
		}
	}
	return types.ImageSummary{}, nil
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

func Start(cli *client.Client, ctx context.Context, id string) error {
	opts := types.ContainerStartOptions{}
	err := cli.ContainerStart(ctx, id, opts)
	if err != nil {
		return err
	}

	// statusCh, errCh := Client.ContainerWait(Ctx, id, container.WaitConditionNotRunning)
	// select {
	// case err := <-errCh:
	//     if err != nil {
	//         panic(err)
	//     }
	// case <-statusCh:

	// }
	// return nil
	return nil
}

func ListImages(cli *client.Client, ctx context.Context) error {
	opts := types.ImageListOptions{
		All: true,
	}
	images, err := cli.ImageList(ctx, opts)
	if err != nil {
		return err
	}

	fmt.Printf("%-50s\t%-6s\t%-8s\t%-25s\t%s\n", "REPOSITORY", "TAG", "IMAGE ID", "CREATED", "SIZE")
	for _, i := range images {
		repotag := s.Split(i.RepoTags[0], ":")
		id := s.Split(i.ID, ":")
		created := time.Unix(i.Created, 0).Format("_2 Jan Mon 2006 03:04:05 PM")
		fmt.Printf("%.50s\t%-6s\t%-8s\t%s\t%.2fMB\n", fmt.Sprintf("%-50sMB", repotag[0]), repotag[1], id[1][:5], created, float64(i.Size)/(1_000_000))
		// fmt.Printf("%v\n", i.RepoTags[1:])
	}
	return nil
}

func GetImage(cli *client.Client, ctx context.Context, nameOrId string) (types.ImageSummary, error) {
	opts := types.ImageListOptions{
		All: true,
	}
	images, err := cli.ImageList(ctx, opts)

	t := types.ImageSummary{}

	if err != nil {
		return t, err
	}
	for _, i := range images {
		if i.ID == nameOrId {
			return i, nil
		}
		for _, name := range i.RepoTags {
			if s.Contains(name, nameOrId) {
				return i, nil
			}
		}
	}
	return t, nil
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

func (d *Manager) GetLatestCachedImgVersion(ctx context.Context, imgName string) (types.ImageSummary, error) {
	images, err := d.Client.ImageList(ctx, types.ImageListOptions{All: true})
	var latestTag string
	var latestImg types.ImageSummary
	if err != nil {
		return latestImg, err
	}

	for _, i := range images {
		for _, tag := range i.RepoTags {
			if s.Contains(tag, imgName) && latestTag < tag {
				fmt.Printf("%v\n", tag)
				latestTag = tag
				latestImg = i
			}
		}
	}

	return latestImg, nil
}

func (d *Manager) Exec(ctx context.Context, containerName string, cmd []string) {
	//containerName := "pokemon-postgres"

	//// Command to execute inside the Docker container
	//cmd := []string{"ls", "-l"}

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
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
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

	return "", fmt.Errorf("tainer not found: %s", containerName)
}
