package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	payload := map[string]interface{}{
		"images": []string{},
	}
	for _, image := range images {
		fmt.Printf("%s %s\n", image.ID[:10], image.ID)
		images, _ := payload["images"].([]string)
		images = append(images, image.ID)
		payload["images"] = images
	}
	result, _ := json.Marshal(payload)
	url := "http://10.10.124.53:8000/api/setContainers"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(result))
	if err != nil {
		fmt.Println("какое то говно", err)
		return
	}
	defer resp.Body.Close()

}
