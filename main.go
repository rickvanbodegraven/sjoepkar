package main

import (
	"fmt"

	// import package under the alias "docker"
	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	// initialize a docker-client connection to the unix socket
	client, _ := docker.NewClientFromEnv()

	// retrieve a list of all docker images
	//	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})

	// retrieve a list of all docker containers
	containers, _ := client.ListContainers(docker.ListContainersOptions{})

	// dump some info for all of these beautiful containers
	for _, container := range containers {
		fmt.Println(container.ID[:12], container.Names[0])

		//		fmt.Println("ID:", container.ID)
		//		fmt.Println("Name:", container.Names)
		//		fmt.Println("Label:", container.Labels)
	}
}
