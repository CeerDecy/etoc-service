package docker

import (
	docker "github.com/fsouza/go-dockerclient"

	"etoc-service/internal/app/http/svc"
	dtype "etoc-service/internal/app/http/types/docker"
	"etoc-service/internal/app/http/types/request"
	"etoc-service/internal/app/http/types/response"
)

func CreateDockerClient(ctx *svc.Context) error {
	var req request.CreateDockerClientRequest
	if err := ctx.ShouldBind(&req); err != nil {
		return err
	}

	//err := docker.CreateDocker(doc)
	return nil
}

func TryConnect(ctx *svc.Context) error {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return err
	}
	version, err := client.Version()
	if err != nil {
		return err
	}
	dockerInfo, err := dtype.ConvertEnvToVersionInfo(version)
	if err != nil {
		return err
	}

	ctx.Success(response.GetLocalDockerConnectionResp{
		EngineVersion: dockerInfo.Version,
		Arch:          dockerInfo.Arch,
		Os:            dockerInfo.Os,
		KernelVersion: dockerInfo.KernelVersion,
		Platform:      dockerInfo.Platform.Name,
	})
	return nil
}
