package dockerScripts

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func PullImageWithAuth() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	authconfig := types.AuthConfig{
		Username: "docker_hub_email",
		Password: "docker_hub_password",
	}
	encodedJSON, err := json.Marshal(authconfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePull(ctx, "httpd", types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		fmt.Println("docker hub account bilgileri eksik olabilir.")
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)

}
