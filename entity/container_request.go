package entity

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
)

// ContainerRequest - A data model to represent the
// Docker Container request entity.
type ContainerRequest struct {
	// Name represents the name that will
	// be assigned to the container.
	// Optional.
	Name string `json:"name"`
	// ImageName represents the name of the
	// image to be pulled.
	// Required value.
	ImageName string `json:"image_name"`
	// Env maps the environment variables.
	// Optional.
	Env map[string]string `json:"env"`
	// Entrypoint - The container entry point.
	// Optional.
	Entrypoint []string `json:"entrypoint"`
}

// Validate - will validate all ContainerRequest's fields.
func (c *ContainerRequest) Validate() error {
	if c.ImageName == "" {
		return fmt.Errorf(`the field '%s' is required`, "image_name")
	}
	return nil
}

// ToContainer will build an instance of Container.
func (c *ContainerRequest) ToContainer() *Container {
	envs := []string{}
	for k, v := range c.Env {
		envs = append(envs, fmt.Sprintf("%s=%s", k, v))
	}
	return &Container{
		Name: c.Name,
		Config: &container.Config{
			Image:      c.ImageName,
			Entrypoint: c.Entrypoint,
			Env:        envs,
		},
	}
}
