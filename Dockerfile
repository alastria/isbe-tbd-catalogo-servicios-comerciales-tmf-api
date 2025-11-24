# Stage 1: Build TMForum API server
FROM golang:1.25.3-bookworm AS builder

# Install build tools for CGO and sqlite tools
# Debian uses apt-get. gcc and libc-dev are usually included in golang images, 
# but we ensure they are there. sqlite3 is useful for tools.
RUN apt-get update && \
    apt-get install -y --no-install-recommends gcc libc6-dev && \
    rm -rf /var/lib/apt/lists/*

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
# Using debian:trixie-slim to ensure glibc >= 2.38 support
FROM debian:trixie-slim

WORKDIR /

COPY --from=builder /isbetmf /isbetmf
COPY www /www
COPY ./auth_policies.star /auth_policies.star
COPY --chmod=755 ./bin/sqlite3_rsync /usr/local/bin/sqlite3_rsync

# Install runtime dependencies
# ca-certificates is required for SSL support
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl wget ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Expose the port the server runs on
EXPOSE 9991

# Run the binary
ENTRYPOINT ["/isbetmf"]
