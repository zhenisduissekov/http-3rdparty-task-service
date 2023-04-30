package service

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	url     string
	method  string
	headers map[string]string
	body    []byte
}

type Task interface {
	TaskQueue() error
	AssignTask(f func())
}

func New() *Service {
	return &Service{}
}

const queueSize = 100

type response struct {
	id             string
	status         string
	httpStatusCode int
	headers        map[string]string
	length         int
	body           []byte //optional
}

var stack = make(map[string]func())
var results = make(map[string]response)
var queue = make(chan string, queueSize)

func (s *Service) TaskQueue() error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case nextTask, ok := <-queue:
			fmt.Println("task received")
			if !ok {
				fmt.Println("channel was closed")
				return nil
			}

			time.Sleep(15 * time.Second)
			for id, f := range nextTask {
				fmt.Println("task", id)
				f()
			}
		case <-ticker.C:
			fmt.Println("Tick", time.Now())
		}
	}

	return nil
}

func (s *Service) AssignTask(url, method string, headers map[string]string, body []byte) (string, error) {
	f := func() {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		client := &http.Client{}
		fmt.Println("requesting-------------->", url)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

	}

	id := uuid.New().String()
	stack[id] = f
	queue <- id

	return id, nil
}

func (s *Service) CheckTask(id string) (string, error) {
	if _, ok := stack[id]; ok {
		return "in progress", nil
	}
	return "done", nil
}
