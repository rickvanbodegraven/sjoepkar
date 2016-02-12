package main

import (
	"fmt"
	"strings"

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
		if strings.Contains(container.Names[0], "cssonline-proxy") {
			fmt.Println(container.Names[0][1:])
			return
		}
	}
}
