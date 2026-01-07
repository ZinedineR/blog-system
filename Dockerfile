FROM golang as builder

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 go build -o blog-system /app/cmd/web/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/blog-system /app/
