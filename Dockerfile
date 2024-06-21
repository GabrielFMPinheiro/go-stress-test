FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o cli .

ENTRYPOINT ["/app/cli"]