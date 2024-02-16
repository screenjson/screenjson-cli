# Use the official Golang image to create a build artifact.
FROM golang:latest as builder

# Create and change to the directory /app.
WORKDIR /app

# First, copy the go.mod and go.sum files from the project root
COPY go.mod go.sum ./

# Fetch any external dependencies that might be required.
RUN go mod download

# Assuming your Go project already uses modules and you have a go.mod file
# Copy the local package files to the container's workspace.
COPY ./src/screenjson-db .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/screenplay-db .

# Use the official Debian slim image for a lean production container.
FROM debian:buster-slim
COPY --from=builder /bin/screenjson-db /screenjson-db

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/screenjson-db"]
