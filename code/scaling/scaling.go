package scaling

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	MemoryThreshold      = 0.1
	MonitoringTimePeriod = 10
)

func monitorServices() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("error creating client: ", err)
		return
	}
	defer cli.Close()
	containers, isSucceed := getAllServices(ctx, cli)
	if isSucceed {
		return
	}

	for _, container := range containers {
		stats, err := cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			fmt.Println("Error getting stats for container ", err)
			continue
		}
		defer stats.Body.Close()

		var statsJSON types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&statsJSON); err != nil {
			fmt.Println("Error decoding stats for container", err)
			continue
		}

		memoryUsage := statsJSON.MemoryStats.Usage
		memoryInGB := convertBytesToGB(memoryUsage)

		isThresholdCrossed := checkThresholdCrossed(memoryInGB)

		if isThresholdCrossed {
			serviceSetUp := ServiceStartUp{}
			serviceSetUp.imageName = container.Image
			if len(container.Ports) > 0 {
				for port, bindings := range container.Ports {
					serviceSetUp.containerPort = strconv.Itoa(port)
					serviceSetUp.hostPort = strconv.Itoa(int(bindings.PublicPort))
				}
			}

			createNewService(serviceSetUp)
		}

	}
}

func checkThresholdCrossed(memoryInGB float64) bool {
	if memoryInGB > MemoryThreshold {
		return true
	}
	return false
}
