# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-socks5-proxy .

# Stage 2: Create the final minimal image
FROM scratch

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/go-socks5-proxy .

# Expose the default SOCKS5 port
EXPOSE 1080

# Run the proxy server
ENTRYPOINT ["/app/go-socks5-proxy"]
