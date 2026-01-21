# Build stage
FROM golang:1.24.1-alpine AS builder

# Add necessary build tools
RUN apk add --no-cache git

# Set Go environment for direct downloads
ENV GOPROXY=direct
ENV GOSUMDB=off

# Set up workspace
WORKDIR /app

# Copy go.work file
COPY go.work ./

# Copy pkg module
COPY pkg ./pkg

# Copy service code
COPY user-service-api ./user-service-api

# Build the application
WORKDIR /app/user-service-api
RUN go mod download && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Create necessary directories
RUN mkdir -p /app/pkg/config && \
    chmod -R 755 /app/pkg

# Copy binary and config
COPY --from=builder /app/user-service-api/main ./main
COPY --from=builder /app/pkg/config/config.yaml ./pkg/config/
COPY .env ./

# Run the application
CMD ["./main"]