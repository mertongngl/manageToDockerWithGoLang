package dockerScripts

import (
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func PullImage() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, "nginx", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
}
