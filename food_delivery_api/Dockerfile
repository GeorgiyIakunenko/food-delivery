# Use the appropriate Go version image as the base image
FROM golang:1.20.5-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only go.mod and go.sum to enable caching for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o fooddelivery

# Specify the entrypoint to run the application
CMD ["./fooddelivery"]

