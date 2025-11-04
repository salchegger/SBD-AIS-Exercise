  # todo
  ## re-used from Exercise 3

# 1: Builder ----------
FROM golang:1.25 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod/go.sum first to leverage Docker layer cache for module downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy everything from your local folder to /app in the container
COPY . .

# Run your build script (creates /app/ordersystem binary)
RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh



# 2: Runner ----------
FROM alpine:latest

# Create a non-root user (recommended)
RUN adduser -D -h /home/appuser appuser

# Install necessary dependencies (for example, ca-certificates)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /app/ordersystem /app/ordersystem

# Copy frontend if the app serves static files (optional)
COPY --from=builder /app/frontend /app/frontend

# Make sure the binary is executable and owned by appuser
RUN chown appuser:appuser /app/ordersystem && chmod +x /app/ordersystem

# Use the non-root user
USER appuser

# Expose application port (change if your app uses a different one)
EXPOSE 3000

# Run the application
ENTRYPOINT ["./ordersystem"]

