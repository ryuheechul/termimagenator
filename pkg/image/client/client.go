package client

import (
	"github.com/docker/docker/client"
)

var singleInstance *client.Client

func Client() (*client.Client, error) {
	if singleInstance != nil {
		return singleInstance, nil
	}

	singleInstance, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	return singleInstance, err
}
