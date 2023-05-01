package service

import (
	"time"
)

const (
	cacheNotInitializedErrMsg = "cache not initialized"
	requestTimeout            = 10 * time.Second
	queueSize                 = 100
	tickPeriod                = 2 * time.Second
	notFoundErrMsg            = "not found"
	statusNew                 = "new"
	statusError               = "error"
	statusInProcess           = "in_process"
	statusDone                = "done"
	cacheLimit                = 1000
	cacheLimitReachedErrMsg   = "cache limit reached, please try again"
	taskReceivedMsg           = "task received"
	channelWasClosedMsg       = "channel was closed"
	tickMsg                   = "tick"
	itemIsNotAssignReq        = "item is not AssignTaskReq"
	failedToPrepareReq        = "failed to prepare request"
	failedRequest             = "request failed"
	failedToCloseRespBody     = "failed to close response body"
	failedToReadRespBody      = "failed to read response body"
	failedToGetItemErrMsg     = "failed to get item"
)

var queue = make(chan string, queueSize)

type AssignTaskReq struct {
	Method  string            `json:"method" validate:"required,min=3,max=6,alphanum,uppercase" example:"GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD, CONNECT, TRACE"`
	Url     string            `json:"url" validate:"required" example:"http://google.com"`
	Headers map[string]string `json:"headers" validate:"omitempty" example:"\"Authentication\": \"Basic bG9naW46cGFzc3dvcmQ=\""`
	ReqBody []byte            `json:"body" validate:"omitempty" example:"{\"name\":\"John\"}"`
	Status  string            `json:"status" validate:"omitempty" example:"done/in_process/error/new"`
}

type AssignTaskResp struct {
	Id             string            `json:"id" example:"4bf1119d-4e7e-4750-99f6-2df3b75acfda"`
	Status         string            `json:"status" example:"done/in_process/error/new"`
	HttpStatusCode int               `json:"httpStatusCode" example:"200"`
	Headers        map[string]string `json:"headers" validate:"omitempty" example:"\"Authentication\": \"Basic bG9naW46cGFzc3dvcmQ=\""`
	Length         int               `json:"length" example:"100"`
	Body           string            `json:"body" example:"{\"name\":\"John\"}"`
}
