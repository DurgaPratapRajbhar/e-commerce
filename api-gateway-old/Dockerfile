# Builder stage
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

 
# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs   

EXPOSE 8081

CMD ["./main"]
