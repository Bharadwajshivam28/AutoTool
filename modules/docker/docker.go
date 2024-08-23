package docker

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func BuildDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	cmd := exec.Command("docker", "build", "-t", dockerImageName, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building Docker image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker build: %v", err)
	}

	fmt.Println("Build complete.")
}

func RunDockerContainer() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)
	containerName := os.Getenv("CONTAINER_NAME")
	hostPort := os.Getenv("HOST_PORT")
	containerPort := os.Getenv("CONTAINER_PORT")

	cmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%s:%s", hostPort, containerPort), "--name", containerName, dockerImageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running Docker container from image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker container: %v", err)
	}

	fmt.Println("Container is running.")
}

func StopDockerContainer() {
	containerName := os.Getenv("CONTAINER_NAME")

	cmd := exec.Command("docker", "stop", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Stopping Docker container:", containerName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error stopping docker container: %v", err)
	}

	fmt.Println("Container stopped.")
}

func RemoveDockerContainer() {
	containerName := os.Getenv("CONTAINER_NAME")

	cmd := exec.Command("docker", "rm", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Removing Docker container:", containerName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error removing docker container: %v", err)
	}

	fmt.Println("Container removed.")
}

func RemoveDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)

	cmd := exec.Command("docker", "rmi", dockerImageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Removing Docker image:", dockerImageName)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error removing docker image: %v", err)
	}

	fmt.Println("Image removed.")
}

func ScanDockerImage() {
	dockerImage := os.Getenv("DOCKER_IMAGE")
	dockerTag := os.Getenv("DOCKER_TAG")
	dockerImageName := fmt.Sprintf("%s:%s", dockerImage, dockerTag)
	cmd := exec.Command("docker", "scout", "quickview", dockerImageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Scanning Docker image:", dockerImageName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running docker scan: %v", err)
	}
	fmt.Println("Scan complete.")
}
