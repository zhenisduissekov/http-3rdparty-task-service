FROM golang:1.19 AS build
RUN mkdir http-3rdparty-task-service && chmod 777 -R ./http-3rdparty-task-service
COPY . /http-3rdparty-task-service
WORKDIR /http-3rdparty-task-service
RUN go test -v ./...
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o .

FROM alpine:latest AS final
RUN mkdir app
RUN chmod 777 -R ./app
COPY --from=build /http-3rdparty-task-service /app
WORKDIR /app
RUN apk add --no-cache tzdata
ENV TZ=Asia/Almaty
CMD ["./http-3rdparty-task-service"]
EXPOSE 3000