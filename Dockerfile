# Dockerfile
FROM golang:1.21.13-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN make build

# Run the application
ENTRYPOINT ["./godeep"]
