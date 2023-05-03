package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
)

func TestGetRespHeaders(t *testing.T) {
	t.Parallel()
	// Create a new HTTP response with test headers
	resp := &http.Response{
		Header: http.Header{
			"Content-Type":  {"application/json"},
			"Cache-Control": {"max-age=3600"},
		},
	}

	// Call the function being tested
	headers := getRespHeaders(resp)

	// Assert that the function returns the correct headers
	expectedHeaders := map[string]string{
		"Content-Type":  "application/json",
		"Cache-Control": "max-age=3600",
	}
	assert.Equal(t, expectedHeaders, headers)
}

func TestPrepReq(t *testing.T) {
	t.Parallel()

	// Define test input
	items := entity.Task{
		Method:  "POST",
		Url:     "https://example.com/api/test",
		ReqBody: `{"foo":"bar"}`,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Basic abc123",
		},
	}

	// Call the function being tested
	req, err := prepReq(items)

	// Check for errors
	if err != nil {
		t.Fatalf("prepReq returned an error: %s", err.Error())
	}

	// Check the request method
	if req.Method != items.Method {
		t.Errorf("prepReq returned a request with the wrong method. Expected %s, got %s", items.Method, req.Method)
	}

	// Check the request URL
	if req.URL.String() != items.Url {
		t.Errorf("prepReq returned a request with the wrong URL. Expected %s, got %s", items.Url, req.URL.String())
	}

	// Check the request body
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("Error reading request body: %s", err.Error())
	}
	if string(reqBody) != items.ReqBody {
		t.Errorf("prepReq returned a request with the wrong body. Expected %s, got %s", items.ReqBody, string(reqBody))
	}

	// Check the request headers
	for k, v := range items.Headers {
		if req.Header.Get(k) != v {
			t.Errorf("prepReq returned a request with the wrong header. Expected %s=%s, got %s=%s", k, v, k, req.Header.Get(k))
		}
	}
}

func TestMakeRequest(t *testing.T) {
	mockClient := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, "OK")
	}))
	defer server.Close()

	task := entity.Task{
		Method: "GET",
		Url:    server.URL,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		ReqBody: "",
	}

	service := &NewService{
		httpClient: mockClient,
	}

	body, headers, status, err := service.makeRequest(task)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(body) != "OK\n" {
		t.Errorf("Unexpected response body: %s", string(body))
	}
	if status != http.StatusOK {
		t.Errorf("Unexpected status code: %d", status)
	}
	expectedHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	if headers["Content-Type"] != expectedHeaders["Content-Type"] {
		t.Errorf("Unexpected headers: got %v, expected %v", headers["Content-Type"], expectedHeaders)
	}
}

func TestNewService_CloseChannel(t *testing.T) {
	t.Cleanup(func() {
		queue = make(chan entity.Task, queueSize)
	})
	cnf := config.New()
	repo := repository.NewRepository(cnf)
	srv := New(repo, cnf)
	srv.CloseQueue()
	if val, ok := <-queue; ok {
		t.Errorf("Channel is not closed: %v", val)
	}
}

func TestNewService_AssignTask(t *testing.T) {
	t.Parallel()
	cnf := config.New()
	repo := repository.NewRepository(cnf)
	srv := New(repo, cnf)
	items := entity.Task{
		Id: "test1",
	}
	id, err := srv.Assign(items)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if id != items.Id {
		t.Errorf("Unexpected id: %v", id)
	}

}

func TestTaskQueue(t *testing.T) {
	t.Cleanup(func() {
		queue = make(chan entity.Task, queueSize)
	})
	cnf := config.New()
	repo := repository.NewRepository(cnf)
	srv := New(repo, cnf)
	go func() {
		srv.StartQueue()
	}()
	queue <- entity.Task{
		Id:      "123",
		Url:     "http://example.com",
		Method:  http.MethodGet,
		ReqBody: "",
	}
	time.Sleep(100 * time.Millisecond)
	if len(queue) != 0 {
		t.Errorf("queue still has items, want empty queue")
	}
	if len(queue) == 0 {
		srv.CloseQueue()
	}
	time.Sleep(100 * time.Millisecond)
}
