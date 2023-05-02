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
if the project was more complicated, then I would use a Makefile to run the project and move main.go to ./cmd folder

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
    swag init -g main.go
```

### Notes

* grouped endpoints with `/api/v1` for future version control

### Used libraries

* I chose fiber v2 for its speed and easy implementation, also it has very good test implementation, if I needed to use as fewer dependencies as possible, then I would use net/http or chi.
However, net/http would require me to write more code.

* Uber fx lets me write less code for dependency injection.

* Zerolog is a very fast logger, it is very easy to use and it has a lot of features. Even though is not updated anymore, it is still a very good logger, therefore does not require updates.
