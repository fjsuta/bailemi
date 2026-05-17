FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/gateway .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/migrations ./migrations

RUN mkdir -p /app/uploads

ENV GIN_MODE=release
ENV SERVER_PORT=8080

EXPOSE 8080

CMD ["./gateway"]
