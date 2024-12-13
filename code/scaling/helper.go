package scaling

import (
	"context"
	"fmt"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ServiceStartUp struct {
	imageName     string
	containerPort string
	hostPort      string
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

func getAllServices(ctx context.Context, cli *client.Client) ([]types.Container, bool) {
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: false})
	if err != nil {
		fmt.Println("Error listing containers:", err)
		return nil, true
	}
	return containers, false
}

func convertBytesToGB(memory uint64) float64 {
	gb := float64(memory) / 1024 / 1024 / 1024
	return gb
}

func createNewService(serviceSetUp ServiceStartUp) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	newContainer, err := setUpService(cli, serviceSetUp.imageName, serviceSetUp.containerPort, serviceSetUp.hostPort)

	if err != nil {
		fmt.Println("Error creating new container", err)
		return
	}

	if err := cli.ContainerStart(ctx, newContainer.ID, container.StartOptions{}); err != nil {
		fmt.Println("Error starting new container", err)
		return
	}

	fmt.Println("Container created and started successfully\n", newContainer.ID)

	// Register the new service by sending port to service discovery so that new service is a part of whole system
}

func setUpService(cli *client.Client, imageName string, containerPort string, hostPort string) (container.CreateResponse, error) {
	newContainer, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		ExposedPorts: map[nat.Port]struct{}{
			nat.Port(containerPort + "/tcp"): {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(containerPort + "/tcp"): []nat.PortBinding{
				{HostPort: hostPort},
			},
		},
	}, nil, nil, "")
	return newContainer, err
}
