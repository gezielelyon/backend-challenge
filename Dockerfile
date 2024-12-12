FROM golang:1.23 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o main .

FROM debian

WORKDIR /app
COPY --from=builder /app/cmd/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]