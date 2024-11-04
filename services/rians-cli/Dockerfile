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

# Use a minimal image for the final output
FROM alpine:latest

# Install any necessary runtime dependencies (optional)
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/rians-cli /usr/local/bin/rians-cli

# Set the entrypoint to the CLI
ENTRYPOINT ["rians-cli"]
