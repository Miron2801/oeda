package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"context"

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
	payload := map[string]interface{}{
		"containers": []string{},
	}
	for _, container := range containers {
		containers, ok := payload["containers"].([]string)
			if ok {
				containers = append(containers, container.ImageID[7:19])
				payload["containers"] = containers
			} else {
				fmt.Println("Список 'containers' не найден")
			}
	}
	result, _ := json.Marshal(payload)
	url := "http://localhost:8080/api/setContainers"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(result))
	if err != nil {
		fmt.Println("какое то говно", err)
		return
	}
	defer resp.Body.Close()

}