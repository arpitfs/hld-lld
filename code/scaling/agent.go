package scaling

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunAgent() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("error creating client: ", err)
		return
	}
	defer cli.Close()

	for {
		containers, err := cli.ContainerList(ctx, container.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing containers: %v\n", err)
			continue
		}

		for _, container := range containers {
			stats, err := cli.ContainerStats(ctx, container.ID, false)
			if err != nil {
				fmt.Printf("Error getting stats for container %s: %v\n", container.ID, err)
				continue
			}
			defer stats.Body.Close()

			var statsJSON types.StatsJSON
			if err := json.NewDecoder(stats.Body).Decode(&statsJSON); err != nil {
				fmt.Printf("Error decoding stats for container %s: %v\n", container.ID, err)
				continue
			}

			// Memory usage stats
			memoryUsage := statsJSON.MemoryStats.Usage
			memoryLimit := statsJSON.MemoryStats.Limit
			fmt.Printf("Container ID: %s, Memory Usage: %d bytes, Memory Limit: %d bytes\n", container.ID, memoryUsage, memoryLimit)
		}

		time.Sleep(10 * time.Second)
	}

}

func createNewService() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	containerName := "my-new-container"
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"sh", "-c", "while true; do echo hello world; sleep 1; done"},
		Tty:   true,
	}, nil, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Printf("Container %s created and started successfully\n", resp.ID)
}

func getAvailablePorts(start, end int) ([]int, error) {
	var availablePorts []int
	for port := start; port <= end; port++ {
		address := fmt.Sprintf("127.0.0.1:%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			availablePorts = append(availablePorts, port)
			listener.Close()
		}
	}
	return availablePorts, nil
}
