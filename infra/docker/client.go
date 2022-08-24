package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
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
func (c *Client) Create(cont *entity.Container) (container.ContainerCreateCreatedBody, error) {
	ctx := context.Background()
	return c.conn.ContainerCreate(
		ctx,
		cont.Config,
		cont.HostConfig,
		cont.NetworkingConfig,
		cont.Platform,
		cont.Name,
	)
}

// Start - Starts a Docker Container
func (c *Client) Start() error { return nil }
