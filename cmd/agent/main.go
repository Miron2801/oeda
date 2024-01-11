package main

import (
	"bytes"
	"context"
	"encoding/json"
	"eoda/intermal/models"
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

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	var images []models.Image
	for _, container := range containers {
		images = append(images, models.Image{Name: container.ImageID, ContainerId: container.ID})
	}
	result, _ := json.Marshal(images)
	url := "http://localhost:8000/api/checkContainers"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(result))
	if err != nil {
		fmt.Println("какое то говно", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var images []models.Image

		err := json.NewDecoder(resp.Body).Decode(&images)

		if err != nil {
			panic(err)
		}
		fmt.Println(images)

		for _, data := range images {
			fmt.Println(data)
			if !data.Accept {
				fmt.Println("kill container", data.ContainerId)
				cli.ContainerKill(context.Background(), data.ContainerId, "9")

			}

		}
	}

	// cli.ContainerKill(context.Background())
}
