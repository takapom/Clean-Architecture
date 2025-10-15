# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install build deps
RUN apk add --no-cache git

# Copy go module manifests first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the bookingapp binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bookingapp ./cmd/api

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/bookingapp ./bookingapp

EXPOSE 8080
ENTRYPOINT ["./bookingapp"]
