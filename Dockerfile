# Use the official Golang image for the build stage
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o blogs-service

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/blogs-service .

# Expose the port the app runs on
EXPOSE 8080

# Define the command to run the application
CMD [ "./blogs-service", "--help" ]
