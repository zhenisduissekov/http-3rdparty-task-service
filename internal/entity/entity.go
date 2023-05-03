package entity

import "time"

const (
	QueueSize                 = 100
	TickPeriod                = 1 * time.Second
	StatusNew                 = "new"
	StatusDone                = "done"
	StatusError               = "error"
	TaskReceivedMsg           = "task received"
	ChannelWasClosedMsg       = "channel was closed, exiting task queue"
	TickMsg                   = "tick"
	FailedToMakeRequestErrMsg = "failed to make request"
)

type Task struct {
	Id             string
	Method         string
	Url            string
	Headers        map[string]string
	ReqBody        string
	RespBody       string
	Status         string
	HttpStatusCode int
	Length         int
}
