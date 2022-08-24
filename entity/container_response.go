package entity

// ContainerResponse - A data model to represent the
// Docker Container response entity.
type ContainerResponse struct {
	// ID is the container ID
	ID string `json:"id"`
	// ImageName represents the name of the
	// image to be pulled.
	// Required value.
	ImageName string `json:"image_name"`
	// Name represents the name that will
	// be assigned to the container.
	// Optional.
	Name string `json:"name"`
	// Status is the current status of the container.
	Status string `json:"status"`
	// Env maps the environment variables.
	// Optional.
	Env map[string]string `json:"env"`
	// Entrypoint - The container entry point.
	// Optional.
	Entrypoint []string `json:"entrypoint"`
}
