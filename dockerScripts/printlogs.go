package dockerScripts

import (
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func PrintLogs() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, "33be5554d9ae", options)

	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
