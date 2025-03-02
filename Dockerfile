# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Install dependencies
RUN apk add --no-cache git curl

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (for better caching)
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the entire source code
COPY . .

# Set working directory to /cmd before building
WORKDIR /app/cmd

# Build the Go application
RUN go build -o /app/main .

# Use a minimal image for production
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy the config.yaml file to the expected location
COPY config/config.yaml /root/config/config.yaml

# Expose the application port
EXPOSE 5000

# Command to run the application
CMD ["./main"]
