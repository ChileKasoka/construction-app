# Stage 1: Build the Go binary
FROM golang:1.24.5-alpine AS builder

WORKDIR /construction-app

# Install git and build tools
RUN apk add --no-cache git build-base

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -x -v -o main .

# Stage 2: Create a minimal container with the binary
FROM alpine:latest

# Set working directory
WORKDIR /construction-app

# Copy the built binary from the builder stage
COPY --from=builder /construction-app/main .

# Expose the application port
EXPOSE 8080

# Set entrypoint
CMD ["./main"]
