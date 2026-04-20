FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/task ./cmd

FROM alpine:3.20
LABEL org.opencontainers.image.source=https://github.com/daniel3579/go_practices_2
WORKDIR /app
COPY --from=builder /app/bin/task /app/task
CMD ["/app/task"]