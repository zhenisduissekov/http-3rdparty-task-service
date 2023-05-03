package repository

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhenisduissekov/http-3rdparty-task-service/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/entity"
)

func TestRepoCache_Set(t *testing.T) {
	cnf := config.New()
	repoCache := NewCache(cnf)
	repoCache.Set("key", entity.Task{
		Id: "id",
	})
	id, err := repoCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "id", id.Id)
}

func TestRepoCache_Set_Overwrite(t *testing.T) {
	cnf := config.New()
	repoCache := NewCache(cnf)
	repoCache.Set("key", entity.Task{
		Id: "id",
	})

	repoCache.Set("key", entity.Task{
		Id: "id2",
	})
	id, err := repoCache.Get("key")
	assert.NoError(t, err)
	assert.Equal(t, "id2", id.Id)
}

func TestRepoCache_Set_Not_Found(t *testing.T) {
	cnf := config.New()
	repoCache := NewCache(cnf)
	id, err := repoCache.Get("key")
	assert.Error(t, err)
	assert.Equal(t, "no task found", err.Error())
	assert.Equal(t, id, entity.Task{})
}

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
	reqBody, err := io.ReadAll(req.Body)
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

	repository := &RepoCache{
		httpClient: mockClient,
	}

	body, headers, status, err := repository.MakeRequest(task)
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
