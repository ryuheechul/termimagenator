package rm

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Remove(images []string) ([]string, []string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	untagged := []string{}
	deleted := []string{}

	if err != nil {
		return untagged, deleted, nil
	}

	for _, image := range images {
		responses, err := cli.ImageRemove(context.Background(), image, types.ImageRemoveOptions{})

		for _, response := range responses {
			if len(response.Untagged) > 0 {
				untagged = append(untagged, response.Untagged)
			}
			if len(response.Deleted) > 0 {
				deleted = append(deleted, response.Deleted)
			}
		}
		if err != nil {
			return untagged, deleted, err
		}
	}
	return untagged, deleted, nil
}
