# --- Build Stage ---
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the API server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /blog-server ./cmd/api
# Build the CLI tool
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /createadmin ./cmd/cli

# --- Final Stage ---
FROM alpine:latest
WORKDIR /
COPY --from=builder /blog-server /blog-server
COPY --from=builder /createadmin /createadmin
EXPOSE 8080
ENTRYPOINT ["/blog-server"]