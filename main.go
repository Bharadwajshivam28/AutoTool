package main

import (
    "fmt"
    "log"
    "os"

    "bharadwajshivam28/autotool/modules/docker"

    "github.com/joho/godotenv"
)

func loadDockerEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func displayHelp() {
    fmt.Println("Ready to help you!")
    fmt.Println("\nDocker Commands:")
    fmt.Println("  bake - Build, scan, and tag a Docker image.")
    fmt.Println("  go - Run a Docker container.")
    fmt.Println("  wipe - Stop and remove Docker containers and images.")
}

func main() {
    if len(os.Args) < 2 {
        displayHelp()
        os.Exit(0)
    }

    command := os.Args[1]
    switch command {
    case "bake":
        loadDockerEnv()
        docker.BuildDockerImage()
        docker.ScanDockerImage()
    case "go":
		loadDockerEnv()
        docker.RunDockerContainer()
    case "wipe":
        loadDockerEnv()
        docker.StopDockerContainer()
        docker.RemoveDockerContainer()
        docker.RemoveDockerImage()
    default:
        fmt.Println("Unknown command:", command)
        displayHelp()
    }
}
