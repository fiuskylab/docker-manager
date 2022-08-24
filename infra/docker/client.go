package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/fiuskylab/docker-api/entity"
	"go.uber.org/zap"
)

// Client manages the Docker API
type Client struct {
	conn *client.Client
}

// NewClient returns a new Instance of Client.
func NewClient() (*Client, error) {
	var err error
	c := new(Client)
	c.conn, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	return c, nil
}

// Conn returns the current instance of Docker client.Client.
func (c *Client) Conn() *client.Client { return c.conn }

// Create - Creates a Docker Container
func (c *Client) Create(cont *entity.Container) error {
	ctx := context.Background()
	body, err := c.conn.ContainerCreate(
		ctx,
		cont.Config,
		cont.HostConfig,
		cont.NetworkingConfig,
		cont.Platform,
		cont.Name,
	)
	if err != nil {
		zap.L().Error(err.Error())
	}

	cont.ID = body.ID

	return err
}

// DeleteContainer - Deletes a container
func (c *Client) DeleteContainer(containerID string) error {
	ctx := context.Background()
	err := c.conn.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		zap.L().Error(err.Error())
	}

	return err
}

// Inspect - will retrieve all data from a container.
func (c *Client) Inspect(containerID string) (types.ContainerJSON, error) {
	ctx := context.Background()
	resp, err := c.conn.ContainerInspect(ctx, containerID)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return resp, err
}

// RetrieveAll return all containers from the
// current running Docker instance.
func (c *Client) RetrieveAll() ([]entity.ContainerResponse, error) {
	list := make([]entity.ContainerResponse, 10)

	ctx := context.Background()
	rawList, err := c.conn.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		zap.L().Error(err.Error())
		return list, err
	}

	for _, l := range rawList {
		name := ""
		if len(l.Names) > 0 {
			name = l.Names[0]
		}
		inspectData, err := c.Inspect(l.ID)
		if err != nil {
			return list, err
		}
		envMap := map[string]string{}
		for _, data := range inspectData.Config.Env {
			envKM := strings.Split(data, "=")
			if len(envKM) == 1 {
				envMap[envKM[0]] = ""
				continue
			}
			envMap[envKM[0]] = envKM[1]
		}
		list = append(list, entity.ContainerResponse{
			Name:       name,
			ImageName:  l.Image,
			Entrypoint: inspectData.Config.Entrypoint,
			Env:        envMap,
			ID:         l.ID,
			Status:     inspectData.State.Status,
		})
	}

	return list, nil
}

// Start - Starts a Docker Container
func (c *Client) Start(cont *entity.Container) error {
	ctx := context.Background()
	if err := c.conn.
		ContainerStart(
			ctx,
			cont.ID,
			types.ContainerStartOptions{},
		); err != nil {
		zap.L().Error(err.Error())
	}

	return nil
}
