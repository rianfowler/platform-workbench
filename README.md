# Project Setup Guide

Welcome to the Project! Follow these instructions to set up your development environment and run the project using Tilt.

## Prerequisites

Ensure the following tools are installed:

1. **Go**
   - Required for building and running the CLI.
   - Installation: [Download and install Go](https://golang.org/dl/).

2. **Pack CLI**
   - Used for building container images.
   - Installation: Follow the [Pack CLI installation guide](https://buildpacks.io/docs/tools/pack/#install).

3. **Score-K8s CLI**
   - Required for generating Kubernetes resources.
   - Installation:
     ```bash
     go install -v github.com/score-spec/score-k8s/cmd/score-k8s@latest
     ```
   - Ensure `GOPATH` is in your `PATH`:
     ```bash
     export PATH=$PATH:$(go env GOPATH)/bin
     ```

4. **Tilt**
   - Used to manage and orchestrate your local development environment.
   - Installation: Follow the [Tilt installation guide](https://docs.tilt.dev/install.html).

5. **Docker**
   - Needed for building and running container images. Ensure Kubernetes support is turned on.
   - Installation: [Download and install Docker Desktop](https://www.docker.com/get-started).
   - Enable Kubernetes:
     - Open Docker Desktop, go to **Settings** > **Kubernetes**, and ensure **Enable Kubernetes** is checked.

6. **Kubectl**
   - Needed for interacting with your local Kubernetes cluster.
   - Installation: Follow the [Kubectl installation guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/).

## Installation Steps

1. **Verify Installations**:
   Ensure all required tools are installed and available in your `PATH`:
   ```bash
   go version        # Verify Go is installed
   pack --version    # Verify Pack CLI is installed
   score-k8s --help  # Verify Score-K8s CLI is installed
   tilt version      # Verify Tilt is installed
   docker version    # Verify Docker is installed
   kubectl version   # Verify Kubectl is installed
   ```

2. **Configure Docker**:
   - Ensure Kubernetes is enabled in Docker Desktop by going to **Settings** > **Kubernetes** and checking **Enable Kubernetes**.

3. **Clone the Repository**:
   - Clone the project repository and navigate to the root directory:
     ```bash
     git clone <repository-url>
     cd <repository-root>
     ```

4. **Run Tilt**:
   - Start the development environment using Tilt:
     ```bash
     tilt up
     ```
   This will start Tilt and initialize the services defined in your Tiltfile.

## Troubleshooting

- **Go Not Found**: Ensure Go is correctly installed and that your `PATH` includes the Go binary.
- **Docker Kubernetes Not Enabled**: Double-check that Kubernetes is enabled in Docker Desktop.
- **Score-K8s CLI Not Found**: Make sure `$(go env GOPATH)/bin` is in your `PATH`.

## Common Commands

- **Rebuild and Restart a Resource Manually**:
   ```bash
   tilt trigger <resource-name>
   ```
- **Stop Tilt**:
   Press `Ctrl+C` in the terminal or stop it from the Tilt UI.

## Additional Notes

- The project includes a `local_resource` for building and installing the CLI. This resource runs manually and does not trigger automatically. You can run it by navigating to the Tilt UI and starting it manually or using:
   ```bash
   tilt trigger install-rians-cli-GO-REQUIRED
   ```

Once `tilt up` is running, you can develop, test, and manage your services in the local environment. Happy coding!