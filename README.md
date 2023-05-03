# http-3rdparty-task-service

## Description

This is an HTTP server that makes requests to 3rd-party services.

* The client sends a task to the service to perform as an http request to 3rd-party services.
* The task is described in json format, the generated task id is returned in response and its execution starts in the background.
* The client must have a method that can be used to find out the status of the task.

## Instructions
* run program
```
    go run ./cmd/main.go
```

* run tests
```
    go test ./... -v
```

* swagger docs
```
    http://localhost:3000/swagger/index.html
```
initializing swagger docs
```
    swag init -g /cmd/main.go

```

useful commands  with docker-compose to operate this service
```
    docker-compose up
    docker-compose up --build
    docker-compose down
```

## Notes

*  endpoints are grouped with `/api/v1` for future version control

## Used libraries

* I chose fiber v2 for its speed and easy implementation, also it has very good test implementation, if I needed to use as fewer dependencies as possible, then I would use net/http or chi.
However, net/http would require me to write more code.

* Zerolog is virtually th fastest logger, it is very easy to use and it has a lot of features. Even though is not updated anymore, it is still a very good logger, therefore does not require updates.

* fiber swagger uses swaggo/swag for swagger documentation, it is very easy and common to use, it has a lot of features and sources.

* testing and assert - common libs for testing.

* To keep data in memory I used go-cache for caching, it is very easy to use and it has a lot of features, but if replacement needed it is implemented in a way of fast replacement.
