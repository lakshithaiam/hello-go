# ---------- Build Stage ----------
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy source
COPY main.go .

# Build binary (static)
RUN go build -o hello-go main.go

# ---------- Runtime Stage ----------
FROM alpine:latest
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/hello-go .

# Expose port and run
EXPOSE 8080
CMD ["./hello-go"]
