# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the server binary specifically
# Change is on this line: ./cmd -> ./cmd/server
RUN CGO_ENABLED=0 go build -o /app/potok-server ./cmd/server

# Stage 2: Create the final, lightweight production image
FROM alpine:latest

WORKDIR /app

# Copy the migrations folder, which is needed at runtime
COPY --from=builder /app/migrations /migrations

# Copy the compiled application binary from the builder stage
COPY --from=builder /app/potok-server .

# Expose the port the application runs on
EXPOSE 8080

# The command to run the application
CMD ["/app/potok-server"]
