package templates

var DockerfileTemplate = `# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/http

# Final stage
FROM alpine:3.18

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Command to run the application
CMD ["./main"]`

func NewDockerfileTemplate() string {
	return DockerfileTemplate
}
