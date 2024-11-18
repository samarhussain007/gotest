# Go Project Setup and Execution Guide

This guide provides a step-by-step explanation of how to initialize, run, and build a simple Go project.

---

## Prerequisites

- **Go Installed**: Ensure that Go is installed on your machine. Download it from [golang.org](https://golang.org/dl/) if you haven’t already.
- **GOPATH Configured**: Verify that your environment is set up correctly with `go env`.

---

## Steps

### 1. Initialize a Go Module

````bash
go mod init hello

This command initializes a new Go module with the name hello. It creates a go.mod file in the current directory, which tracks dependencies for your project.

### 2. Run the Go Program
```bash
go run .

This command compiles and runs the Go program located in the current directory. Ensure you have a main.go file in the directory, as this file is the entry point for the Go program.

### 3. Build the Go Program
```bash
go build .


This command compiles the Go source code in the current directory and creates an executable binary. The resulting binary will have the same name as the module (hello) and can be run independently without the Go runtime.
````
