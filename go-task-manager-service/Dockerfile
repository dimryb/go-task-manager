# Step 1: Caching Go modules
FROM golang:alpine AS modules
WORKDIR /modules
COPY go.mod go.sum ./
RUN go mod download

# Step 2: Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY --from=modules /go/pkg /go/pkg
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/app

# Step 3: Final minimal image
FROM scratch
COPY --from=builder /bin/app /app
COPY --from=builder /app/config /config
COPY --from=builder /app/migrations /migrations
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/go.mod /app/go.sum /
COPY .env /
CMD ["/app"]