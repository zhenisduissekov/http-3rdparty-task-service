# http-3rdparty-task-service

## Description

This is an HTTP server that makes requests to 3rd-party services.

* The client sends a task to the service to perform as an http request to 3rd-party services.
* The task is described in json format, the generated task id is returned in response and
its execution starts in the background.

### Instructions
* run program
```
    go run main.go
```
if needed to specify port
```
    go run main.go --port :4000
```
if the project was more complicated, then I would use a Makefile to run the project and move main.go to ./cmd folder to run it like go run ./cmd/main.go

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

to run with docker-compose
```
    docker-compose up
```

to stop docker-compose
```
    docker-compose down
```

### Notes
swag init -d ./cmd
* grouped endpoints with `/api/v1` for future version control

### Used libraries

* I chose fiber v2 for its speed and easy implementation, also it has very good test implementation, if I needed to use as fewer dependencies as possible, then I would use net/http or chi.
However, net/http would require me to write more code.

* Zerolog is a very fast logger, it is very easy to use and it has a lot of features. Even though is not updated anymore, it is still a very good logger, therefore does not require updates.

* I used swag for swagger documentation, it is very easy to use and it has a lot of features.

* I used testify for testing, it is very easy to use and it has a lot of features.

* I used go cache for caching, it is very easy to use and it has a lot of features, but also using repository makes it easy to replace my choice.