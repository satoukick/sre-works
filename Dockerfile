# ---- Build Stage ----
FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o sre-works .

# ---- Runtime Stage ----
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/sre-works .

EXPOSE 8080

CMD ["./sre-works"]
