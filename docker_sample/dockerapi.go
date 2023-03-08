package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

var (
	//如果是远程链接docker，记得关闭机器的防火墙和修改配置，开启远程访问。
	cli, err = client.NewClientWithOpts(client.WithHost("tcp://192.168.181.128:2375"))

	//链接本地docker的话用这个
	//cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
)

func main() {
	ctx := context.Background()
	images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}
	list, _ := json.Marshal(images)
	fmt.Println(string(list))
	fmt.Println("image size:", len(images))
}
