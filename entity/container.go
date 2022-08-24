package entity

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// Container - A data model to store all DockerContainer
// necessary data.
type Container struct {
	Config           *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
	Platform         *v1.Platform
	Name             string
	ID               string
}
