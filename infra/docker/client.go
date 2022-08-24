package docker

import (
	"context"

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

// Inspect - will retrieve all data from a container.
func (c *Client) Inspect(containerID string) (types.ContainerJSON, error) {
	ctx := context.Background()
	resp, err := c.conn.ContainerInspect(ctx, containerID)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return resp, err
}
