# Use the official Golang image as a build stage
FROM golang:1.22.6-alpine AS builder

# Set the working directory
WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache git gcc g++ libc-dev

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o /loan-service

# Build the application

# Use a minimal image for the final stage
FROM alpine:latest

# Install required system dependencies
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /root/

COPY .env .

# Copy the binary from the build stage
COPY --from=builder /loan-service .

# Expose the application on port 8080
EXPOSE 8080

# Run the application
CMD ["./loan-service"]
