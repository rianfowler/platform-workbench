# Start with the official Go image as the base
FROM golang:1.23 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the CLI binary with the correct name
RUN go build -o rians-cli main.go

# Final stage using a Debian-based image for compatibility
FROM debian:latest
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/rians-cli /usr/local/bin/rians-cli
WORKDIR /workspace
ENTRYPOINT ["rians-cli"]