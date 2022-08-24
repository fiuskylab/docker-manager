package entity

import "fmt"

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
