# --- Build Stage ---
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /blog-server ./cmd/api

# --- Final Stage ---
FROM alpine:latest
WORKDIR /
COPY --from=builder /blog-server /blog-server
EXPOSE 8080
ENTRYPOINT ["/blog-server"]