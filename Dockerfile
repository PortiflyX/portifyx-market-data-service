FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY portifyx-goutils /portifyx-goutils
COPY portifyx-market-data-service /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/market-data main.go

FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/bin/market-data /app/market-data
COPY portifyx-market-data-service/.env.example /app/.env

EXPOSE 8083
CMD ["/app/market-data", "api"]
