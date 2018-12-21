package dockerScripts

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func ListAllImages() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	ctx := context.Background()
	if err != nil {
		panic(err)
	}
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}
