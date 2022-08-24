package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/fiuskylab/docker-api/entity"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	// should not run if log is not set
	if err != nil {
		panic(err)
	}

	// replacing global logger, is not recommended
	// but for this project I'll do it to avoid
	// dependencies management related to log.
	undo := zap.ReplaceGlobals(logger)

	defer undo()

	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	cont := entity.ContainerRequest{
		ImageName: "hello-world",
	}

	out, err := dockerClient.ImagePull(ctx, cont.ImageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)

	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image:      cont.ImageName,
		Entrypoint: cont.Entrypoint,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err = dockerClient.
		ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
