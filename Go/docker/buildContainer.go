package docker

import (
	"context"
	"fmt"
)

func BuildContainer(ctx context.Context, build ContainerBuild) error{
	err := CheckColima()
	if err != nil {
		return fmt.Errorf("colima instance not started -> %w", err)
	}
	err = ValidateBuild(build)
	if err != nil {
		return err
	}


	d := Manager{}
	err = d.Connect(ctx)
	if err != nil {
		return fmt.Errorf("fatal Connect() error -> %w", err)
	}
	defer d.Client.Close()

	image, err := d.FindCachedImage(ctx, fmt.Sprintf("%s:%s", build.ImgName, build.Version))
	if err != nil {
		return fmt.Errorf("find cached image failed, imagelist err: -> %w", err)
	}

	if image.ID == "" { //no cached image time to pull one from registry
		//todo prevent SYSERR if no connection to registry

		switch build.Version {
		case "":
			err = d.Pull(ctx, build.ImgName)
		default:
			err = d.Pull(ctx, fmt.Sprintf("%s:%s", build.ImgName, build.Version))
		}

		if err != nil {
			return fmt.Errorf("pull failed -> %w", err)
		}
	}

	if _, err = d.GetContainer(ctx, build.ContainerName); err == nil{
		err = d.DeleteContainer(ctx, build.ContainerName)
		if err != nil {
			return fmt.Errorf("failed to delete and stop container -> %w", err)
		}
	}
	_ = d.FreeUsedPort(ctx, build.Port)

	build.ContainerId, err = d.RunContainer(ctx, build)
	if err != nil {
		return fmt.Errorf("run container failed -> %w", err)
	}

	return nil
}