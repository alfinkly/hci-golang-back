# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o pharmacy-api main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/pharmacy-api .
COPY --from=builder /app/.env.example .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./pharmacy-api"]
