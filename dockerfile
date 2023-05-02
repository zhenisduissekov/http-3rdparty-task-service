FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o app ./cmd

EXPOSE 3000

CMD ["./app"]
