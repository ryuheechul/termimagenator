package ls

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/ryuheechul/termimagenator/pkg/image"
	"github.com/ryuheechul/termimagenator/pkg/image/formatter"
	"github.com/samber/lo"
)

func List() ([]image.Image, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return []image.Image{}, err
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return []image.Image{}, err
	}
	return lo.FlatMap(images, func(img types.ImageSummary, index int) []image.Image {
		return lo.Map(img.RepoTags, func(repoTag string, index int) image.Image { return image.Image{RepoTag: repoTag, ID: img.ID} })
	}), nil
}

func ListFormatted(f image.Formatter) ([]string, error) {
	images, err := List()

	if err != nil {
		return []string{}, err
	}

	var errCatched error

	return lo.Map(images, func(img image.Image, index int) string {
		formatted, err := f.Format(img)
		if err != nil {
			errCatched = err
			return ""
		}
		return formatted
	}), errCatched
}

func ListWithDefaultFormat() ([]string, error) {
	return ListFormatted(formatter.DefaultFormatter)
}
