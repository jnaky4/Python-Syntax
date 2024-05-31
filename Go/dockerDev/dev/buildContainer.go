package dev

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
)

//todo should these return the gracefulShutdown? build network definitely should since network needs cleaned up

func BuildContainer(ctx context.Context, build ContainerBuild) (Manager, error) {
	d := Manager{}
	err := CheckColima()
	if err != nil {
		return d, fmt.Errorf("colima instance not started -> %w", err)
	}

	err = ValidateBuild(build)
	if err != nil {
		return d, err
	}

	err = d.ClientConnect() //
	if err != nil {
		return d, fmt.Errorf("fatal ClientConnect() error -> %w", err)
	}

	defer d.Client.Close()

	image, err := d.FindCachedImage(ctx, build.Config.Image)
	if err != nil {
		return d, fmt.Errorf("find cached image failed, imagelist err: -> %w", err)
	}

	if image.ID == "" { //if ID is empty, there's no cached image. time to pull one from registry

		err = d.Pull(ctx, build.Config.Image)

		if err != nil {
			return d, fmt.Errorf("pull failed -> %w", err)
		}
	}

	if _, err = d.GetContainer(ctx, build.ContainerName); err == nil {
		err = d.DeleteContainer(ctx, build.ContainerName)
		if err != nil {
			fmt.Printf("failed to delete and stop container -> %s\n", err.Error())
		}
	}

	for k, _ := range build.HostConfig.PortBindings {
		_ = d.FreeUsedPort(ctx, k.Port())
	}

	build.ContainerId, err = d.RunContainer(ctx, build)
	if err != nil {
		return d, fmt.Errorf("run container failed -> %w", err)
	}

	return d, nil
}

// BuildNetwork create a network and returns a GracefulShutdown for the network
func BuildNetwork(ctx context.Context, networkName string, networkBuild types.NetworkCreate) (Manager, error) {
	d := Manager{}

	err := CheckColima()
	if err != nil {
		return d, fmt.Errorf("colima instance not started -> %w", err)
	}

	err = d.ClientConnect() //
	if err != nil {
		return d, fmt.Errorf("fatal ClientConnect() error -> %w", err)
	}

	defer d.Client.Close()

	_, err = d.CleanCreateNetwork(ctx, networkName, networkBuild)
	if err != nil {
		return d, err
	}
	return d, nil

}
