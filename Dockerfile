# Build stage
CMD ["./main"]
# Run the application

EXPOSE 8082
# Expose the application port

USER appuser
# Use non-root user

RUN chown -R appuser:appuser /app
# Change ownership to non-root user

COPY .env .env
# Copy .env file (optional, can be overridden by environment variables)

COPY --from=builder /app/migrations ./migrations
# Copy migrations

COPY --from=builder /app/main .
# Copy the binary from builder

WORKDIR /app

    adduser -D -u 1000 -G appuser appuser
RUN addgroup -g 1000 appuser && \
# Create a non-root user

RUN apk --no-cache add ca-certificates
# Install ca-certificates for HTTPS requests

FROM alpine:latest
# Final stage

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api
# Build the application

COPY . .
# Copy the source code

RUN go mod download
# Download dependencies

COPY go.mod go.sum ./
# Copy go mod files

WORKDIR /app
# Set working directory

RUN apk add --no-cache git ca-certificates tzdata
# Install git and ca-certificates (needed for fetching dependencies)

FROM golang:1.24-alpine AS builder

