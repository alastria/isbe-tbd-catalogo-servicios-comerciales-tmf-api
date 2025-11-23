# Stage 1: Build sqlite3_rsync
FROM alpine:3.20 AS rsyncbuilder

# Install build dependencies
RUN apk add --no-cache build-base wget unzip

# Set SQLite version (check https://sqlite.org/download.html for latest)
ENV SQLITE_VERSION=3510000

# Download SQLite source (amalgamation form)
WORKDIR /tmp
RUN wget https://sqlite.org/2025/sqlite-src-${SQLITE_VERSION}.zip \
    && unzip sqlite-src-${SQLITE_VERSION}.zip

# Build sqlite3_rsync
WORKDIR /tmp/sqlite-src-${SQLITE_VERSION}/tool
RUN gcc -O2 -o sqlite3_rsync sqlite3_rsync.c ../sqlite3.c -lpthread -ldl

# Stage 2: Build TMForum API server
FROM golang:1.25.3-alpine AS builder

# Install build tools for CGO
RUN apk add --no-cache gcc musl-dev

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
COPY --from=rsyncbuilder --chmod=755 /tmp/sqlite-src-${SQLITE_VERSION}/tool/sqlite3_rsync /usr/local/bin/sqlite3_rsync

# Expose the port the server runs on
EXPOSE 9991

# Run the binary
ENTRYPOINT ["/isbetmf"]