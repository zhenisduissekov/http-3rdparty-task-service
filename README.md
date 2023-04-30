# http-3rdparty-task-service

## Description

This is an HTTP server that makes requests to 3rd-party services.

* The client sends a task to the service to perform as an http request to 3rd-party services.
* The task is described in json format, the generated task id is returned in response and
its execution starts in the background.

## Request example to service:

Request:
POST /task
```
{
    "method": "GET",
    "url": "http://google.com",
    "headers": {     "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
    ....
   }
}
```
Response:
200 OK

```
{
    "id": <generated unique id>
}
```

* The client must have a method that can be used to find out the status of the task.

Request example
Request:
```
GET task/<taskId>
```

Response:
200 OK
```
{
    "id"
    :{
   },
}
```

We'd like to see code close to production with clear variable names and http routes, unit tests, etc.
    : <unique id>,
 "status"
"headers"
:
"done/in_process/error/new"
  "httpStatusCode"
: <HTTP status of 3rd-party service response>,

<headers array from 3rd-party service response>
  "length"
 : <content length of 3rd-party service response>


```
http-3rdparty-task-service/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   └── handler.go
│   └── service/
│       └── service.go
├── pkg/
│   ├── database/
│   │   └── database.go
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```
### Instructions
* run program
```
    go run ./cmd/main.go
```
if needed to specify port
```
    go run ./cmd/main.go --port :4000
```
* swagger
```
swag init -g cmd/main.go
```

### Notes

* added /api/v1 for endpoints for future versioning

### Used libraries

I chose fiber v2 for its speed and easy implementation, if I needed to use as less dependencies as possible then I would use net/http or chi.
However, net/http would require me to write more code.

Uber fx lets me write less code for dependency injection.

PgxPool is very good with connection pooling and it is very fast, however a bit tricky for mocking.

Zerolog is a very fast logger, it is very easy to use and it has a lot of features. Even though is not updated anymore, it is still a very good logger, therefore does not require updates.
