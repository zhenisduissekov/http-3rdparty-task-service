FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o app ./cmd

RUN go test -v ./...

EXPOSE 3000

CMD ["./app"]
