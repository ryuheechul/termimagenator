package ls

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/ryuheechul/termimagenator/pkg/image"
	"github.com/ryuheechul/termimagenator/pkg/image/formatter"
	"github.com/samber/lo"
)

func List() []image.Image {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	return lo.FlatMap(images, func(img types.ImageSummary, index int) []image.Image {
		return lo.Map(img.RepoTags, func(repoTag string, index int) image.Image {
			return image.Image{RepoTag: repoTag, ID: img.ID}
		})
	})
}

func ListFormatted(f image.Formatter) []string {
	return lo.Map(List(), func(img image.Image, index int) string {
		formatted, err := f.Format(img)
		if err != nil {
			panic(err)
		}
		return formatted
	})
}

func ListWithDefaultFormat() []string {
	return ListFormatted(formatter.DefaultFormatter)
}
