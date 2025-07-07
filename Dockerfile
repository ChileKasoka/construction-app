# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /construction-app

# Install git and other dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create a lightweight container
FROM alpine:latest

WORKDIR /construction-app

# Copy the binary from the builder
COPY --from=builder /construction-app/main .

# Expose application port
EXPOSE 8080

# Run the binary
CMD ["./main"]
