# Stage 1: Build TMForum API server
FROM golang:1.25.3-alpine AS builder

# Install build tools for CGO and sqlite tools
RUN apk update && \
    apk add --no-cache gcc musl-dev sqlite-tools && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary with CGO enabled
# -ldflags="-w -s" strips debug information and symbols, reducing the binary size
RUN go build -ldflags="-w -s" -o /isbetmf .

# Final stage
FROM alpine/curl:latest

WORKDIR /
COPY --from=builder /isbetmf /isbetmf
COPY www /www
COPY ./auth_policies.star /auth_policies.star
COPY --from=builder --chmod=755 /usr/bin/sqlite3_rsync /usr/local/bin/sqlite3_rsync

# Expose the port the server runs on
EXPOSE 9991

# Run the binary
ENTRYPOINT ["/isbetmf"]