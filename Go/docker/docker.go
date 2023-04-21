package docker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	s "strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	_ "github.com/docker/docker/pkg/stdcopy"
	_ "github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	erp "github.com/pkg/errors"
)

// var ctx context.Context
// var cli client.Client

func main() {
	err := StartColima()
	if err != nil {
		fmt.Printf("Colima Error: %s\n", erp.Cause(err).Error())
		os.Exit(1)
	}

	cli, ctx, err := Connect()
	if err != nil {
		fmt.Printf("Fatal Startup Error -> Cause: %s\n", erp.Cause(err).Error())
		os.Exit(1)
	}
	defer cli.Close()

	// Pull(cli, ctx, "postgres")
	// ListImages(cli, ctx)
	// i, err := GetImage(cli,ctx, "postgres")
	// if err != nil{
	// 	fmt.Printf("%+v\n", erp.Cause(err))
	// }
	// fmt.Printf("Image: %+v\n", i)

	err = PostgresContainer(cli, ctx)
	if err != nil {
		println("err: ", erp.Cause(err).Error())
	}

	println("Waiting 20s")
	time.Sleep(time.Duration(time.Second) * 20)
	postgres, err := GetContainer(cli, ctx, "postgres")
	if err != nil {
		println("err: ", erp.Cause(err).Error())
	}
	err = DeleteContainer(cli, ctx, postgres.ID)
	if err != nil {
		println("err: ", erp.Cause(err).Error())
	}

	// PruneContainer(cli, ctx)
	// ListContainers(cli,ctx, true)
	// err = stop(cli, ctx, id)
	// if err != nil{
	// 	println("err: ", erp.Cause(err))
	// }

	// err = DeleteContainer(cli, ctx, id)
	// if err != nil{
	// 	println("err: ", erp.Cause(err))
	// }

	// out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	// if err != nil {
	//     panic(err)
	// }

	// stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

func StartColima() error {
	err := exec.Command("colima", "status").Run()
	if err != nil {
		//println(erp.Cause(err).Error())
		cmd := exec.Command("colima", "Start")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return erp.Wrap(err, "Failed to Start Colima\n")
		}
	}
	return nil
}

func GetUser() (string, error) {
	files, err := ioutil.ReadDir("/Users/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if s.HasPrefix(f.Name(), "Z") && len(f.Name()) == 7 {
			return f.Name(), nil
		}
	}
	return "", erp.New("cannot find User zID folder in system")
}

func ListContainers(cli *client.Client, ctx context.Context, all bool) error {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: all})
	if err != nil {
		return erp.Wrap(err, "failed to get containerList")
	}

	fmt.Printf("%-12s\t%-30s\t%-10s\t%s\n", "CONTAINER ID", "IMAGE", "STATUS", "PORTS")
	for _, c := range containers {
		fmt.Printf("%-12s\t%.30s\t%.10s\t%v\n", c.ID[:12], fmt.Sprintf("%-30s", c.Image), c.Status, c.Ports)
	}
	return nil
}

func GetContainer(cli *client.Client, ctx context.Context, nameOrId string) (types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	ct := types.Container{}
	if err != nil {
		return ct, erp.Wrap(err, "failed to get containerList")
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
	return ct, nil
}

func Connect() (*client.Client, context.Context, error) {
	if runtime.GOOS == "darwin" {
		uID, err := GetUser()
		if err != nil {
			log.Fatal(err)
		}
		sockPath := fmt.Sprintf("unix:///Users/%s/.colima/default/docker.sock", uID)
		err = os.Setenv("DOCKER_HOST", sockPath)
		if err != nil {
			return nil, nil, erp.Wrap(err, fmt.Sprintf("Failed to find Docker Socket at path %s", sockPath))
		}
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, nil, err
	}
	return cli, ctx, nil
}

func Pull(cli *client.Client, ctx context.Context, image string) error {

	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	defer reader.Close()
	if err != nil {
		return erp.Wrap(err, fmt.Sprintf("Failed to Pull image %s", image))
	}
	io.Copy(os.Stdout, reader)

	return nil
}

func PruneContainer(cli *client.Client, ctx context.Context) {
	fil := filters.Args{}
	cli.ContainersPrune(ctx, fil)
}

func DeleteContainer(cli *client.Client, ctx context.Context, id string) error {
	opts := types.ContainerRemoveOptions{RemoveVolumes: true, Force: true}
	err := cli.ContainerRemove(ctx, id, opts)
	if err != nil {
		return erp.Wrap(err, fmt.Sprintf("Failed to DeleteContainer container %s", id))
	}
	return nil
}

//func stop(cli *client.Client, ctx context.Context, id string) error {
//	err := cli.ContainerStop(ctx, id, nil)
//	if err != nil {
//		return erp.Wrap(err, "Failed to stop container")
//	}
//	return nil
//}

func Start(cli *client.Client, ctx context.Context, id string) error {
	opts := types.ContainerStartOptions{}
	err := cli.ContainerStart(ctx, id, opts)
	if err != nil {
		return erp.Wrap(err, "Failed to Start container")
	}

	// statusCh, errCh := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
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

func Create(cli *client.Client, ctx context.Context, tainerConfig *container.Config, hostCon *container.HostConfig,
	netConfig *network.NetworkingConfig, platform *v1.Platform, cName string) (string, error) {
	cont, err := cli.ContainerCreate(ctx, tainerConfig, hostCon, netConfig, platform, cName)
	if err != nil {
		return "", erp.Wrap(err, "Failed to Create Container")
	}
	return cont.ID, nil
}

// func PostgresContainer(cli *client.Client, ctx context.Context) (tainerConfig *container.Config, hostCon *container.HostConfig,
//
//	netConfig *network.NetworkingConfig, platform *v1.Platform, cName string, err error){
//
// func PostgresContainer(cli *client.Client, ctx context.Context) (*container.Config, *container.HostConfig,
// *network.NetworkingConfig, *v1.Platform, string, error){
func PostgresContainer(cli *client.Client, ctx context.Context) error {
	// expPort := "5000"
	cName := "postgres"
	// container = client.containers.run(
	//     image_name,
	//     detach=True,
	//     name=container_name,
	//     ports={'5000/tcp': 5000}
	//     # environment={"POSTGRES_PASSWORD": "pokemon", "POSTGRES_DB": "Pokemon"}
	// )

	// newport, err := nat.NewPort("tcp", expPort)
	// if err != nil{
	// 	erp.Wrap(err, fmt.Sprintf("Unable to Create docker port: %s", newport.Port()))
	// }

	// exposedPorts := map[nat.Port]struct{}{
	// 	newport: struct{}{},
	// }

	tainerConfig := container.Config{
		Image: "postgres",
		Env:   []string{"POSTGRES_PASSWORD=password"},
		// ExposedPorts: exposedPorts,//List of exposed ports

		// AttachStdin: true, //Attach the standard input, makes possible user interaction

		// Cmd:   []string{"echo", "hello world"},
		// WorkingDir: "/",
	}

	// *hostCon = container.HostConfig{
	// 	PortBindings: nat.PortMap{
	// 		newport: []nat.PortBinding{
	// 			{
	// 				HostIP:   "0.0.0.0",
	// 				HostPort: expPort,
	// 			},
	// 		},
	// 	},
	// 	RestartPolicy: container.RestartPolicy{
	// 		Name: "always",
	// 	},
	// 	LogConfig: container.LogConfig{
	// 		Type:   "json-file",
	// 		Config: map[string]string{},
	// 	},
	// }

	// *netConfig = network.NetworkingConfig{
	// 	EndpointsConfig: map[string]*network.EndpointSettings{},
	// }

	// platform = &v1.Platform{}

	// return &tainerConfig, nil, nil, nil, cName, nil

	fmt.Printf("Creating Container %s\n", cName)
	id, err := Create(cli, ctx, &tainerConfig, nil, nil, nil, cName)
	if err != nil {
		return err
	}
	fmt.Printf("Starting Container %s, id: %s\n", cName, id[:5])
	err = Start(cli, ctx, id)
	if err != nil {
		return err
	}
	return nil
}
//type ExecResult struct {
//	StdOut string
//	StdErr string
//	ExitCode int
//}

//func PostgresReady(cli *client.Client, ctx context.Context, cID string) (ExecResult,error) {
//	var execResult ExecResult
//	//ecfg := types.ExecConfig{
//	//	AttachStderr: true,
//	//	AttachStdout: true,
//	//	Cmd: []string{"pg_isready"},
//	//}
//	//create, err := cli.ContainerExecCreate(ctx, cID, ecfg)
//	//if err != nil {
//	//	return execResult, err
//	//}
//	//err = cli.ContainerExecStart(ctx, create.ID, types.ExecStartCheck{})
//	//if err != nil {
//	//	return ExecResult{}, err
//	//}
//
//	//return execResult, nil
//
//	resp, err := cli.ContainerExecAttach(ctx, cID, types.ExecStartCheck{})
//	if err != nil {
//		return execResult, err
//	}
//	defer resp.Close()
//
//	//read the output
//	var outBuf, errBuf bytes.Buffer
//	outputDone := make(chan error)
//
//	go func() {
//		// StdCopy demultiplexes the stream into two buffers
//		_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
//		outputDone <- err
//	}()
//
//	select {
//	case err := <-outputDone:
//		if err != nil {
//			return execResult, err
//		}
//		break
//
//	case <-ctx.Done():
//		return execResult, ctx.Err()
//	}
//
//	stdout, err := io.ReadAll(&outBuf)
//	if err != nil {
//		return execResult, err
//	}
//	stderr, err := io.ReadAll(&errBuf)
//	if err != nil {
//		return execResult, err
//	}
//
//	res, err := cli.ContainerExecInspect(ctx, cID)
//	if err != nil {
//		return execResult, err
//	}
//
//	execResult.ExitCode = res.ExitCode
//	execResult.StdOut = string(stdout)
//	execResult.StdErr = string(stderr)
//	return execResult, nil
//
//}
type ExecResult struct {
	ExitCode  int
	outBuffer *bytes.Buffer
	errBuffer *bytes.Buffer
}

func Exec(ctx context.Context, cli client.APIClient, id string, cmd []string) (ExecResult, error) {
	// prepare exec
	execConfig := types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
	}
	cresp, err := cli.ContainerExecCreate(ctx, id, execConfig)
	if err != nil {
		return ExecResult{}, err
	}
	execID := cresp.ID

	// run it, with stdout/stderr attached
	aresp, err := cli.ContainerExecAttach(ctx, execID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}
	defer aresp.Close()

	// read the output
	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		// StdCopy demultiplexes the stream into two buffers
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, aresp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return ExecResult{}, err
		}
		break

	case <-ctx.Done():
		return ExecResult{}, ctx.Err()
	}

	// get the exit code
	iresp, err := cli.ContainerExecInspect(ctx, execID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{ExitCode: iresp.ExitCode, outBuffer: &outBuf, errBuffer: &errBuf}, nil
}
// Stdout returns stdout output of a command run by Exec()
func (res *ExecResult) Stdout() string {
	return res.outBuffer.String()
}

// Stderr returns stderr output of a command run by Exec()
func (res *ExecResult) Stderr() string {
	return res.errBuffer.String()
}

// Combined returns combined stdout and stderr output of a command run by Exec()
func (res *ExecResult) Combined() string {
	return res.outBuffer.String() + res.errBuffer.String()
}