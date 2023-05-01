FROM golang:1.20 AS build
RUN mkdir partners && chmod 777 -R ./partners
COPY . /partners
WORKDIR /partners
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o .


FROM alpine:latest AS final
RUN mkdir app
RUN chmod 777 -R ./app
COPY --from=build /partners /app
WORKDIR /app
RUN apk add --no-ccache tzdata
ENV TZ=Asia/Almaty
CMD ["./partners"]
EXPOSE 3000