package docker

import (
	"github.com/docker/docker/client"
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
