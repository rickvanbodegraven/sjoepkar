package main

import (
	"os"
    "fmt"

    "github.com/fsouza/go-dockerclient"
)

func main() {
	// retrieve the $DOCKER_HOST environment variable 
    endpoint := os.Environ()["DOCKER_HOST"]
	
    client, _ := docker.NewClient(endpoint)
	
    imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	
    for _, img := range imgs {
        fmt.Println("ID: ", img.ID)
        fmt.Println("RepoTags: ", img.RepoTags)
        fmt.Println("Created: ", img.Created)
        fmt.Println("Size: ", img.Size)
        fmt.Println("VirtualSize: ", img.VirtualSize)
        fmt.Println("ParentId: ", img.ParentID)
    }
}