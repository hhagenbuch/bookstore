# Use the official Golang image to create a build artifact.
# This is based on the alpine image, which is a small Linux distribution.
FROM golang:1.22.4-alpine AS builder

# Install dependencies for CGO
RUN apk update && apk add --no-cache gcc musl-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with CGO enabled
ENV CGO_ENABLED=1
RUN go build -o bookstore cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Install SQLite
RUN apk --no-cache add sqlite

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bookstore .

# Copy the SQLite database file
COPY bookstore.db .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./bookstore"]
