# Use the official Golang image to create a build artifact.
FROM golang:latest as builder

# Create and change to the directory /app.
WORKDIR /app

# Assuming your Go project already uses modules and you have a go.mod file
# Copy the local package files to the container's workspace.
COPY ./src/screenjson-db .

# Fetch any external dependencies that might be required.
RUN go mod download

# Build the command inside the container.
RUN go build -o /screenjson-db

# Use the official Debian slim image for a lean production container.
FROM debian:buster-slim
COPY --from=builder /screenjson-db /screenjson-db

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/screenjson-db"]
