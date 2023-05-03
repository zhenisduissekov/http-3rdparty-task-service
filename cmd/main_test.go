package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/zhenisduissekov/http-3rdparty-task-service/config"
)

func Test_appHealth(t *testing.T) {
	app := router(nil, &config.Conf{ServiceName: "test"})

	type args struct {
		method   string
		endpoint string
		body     []byte
	}
	type want struct {
		status    int
		message   string
		checkBody bool
	}
	tests := []struct {
		name string
		desc string
		args args
		want want
	}{
		{
			name: "test1",
			desc: "check health responds with 200",
			args: args{
				method:   "GET",
				endpoint: "/health",
				body:     nil,
			},
			want: want{
				status:    http.StatusOK,
				message:   "{\"status\":\"success\"}",
				checkBody: true,
			},
		},
		{
			name: "test2",
			desc: "check swagger responds with 200",
			args: args{
				method:   "GET",
				endpoint: "/swagger/index.html",
				body:     nil,
			},
			want: want{
				status:    http.StatusOK,
				message:   "",
				checkBody: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(tt.args.method, tt.args.endpoint, bytes.NewBuffer(tt.args.body))
			if err != nil {
				t.Fatal(err)
			}

			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != tt.want.status {
				t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
			}

			if tt.want.checkBody {
				actual, _ := io.ReadAll(resp.Body)
				if string(actual) != tt.want.message {
					t.Errorf("Expected response body %q but got %q", tt.want.message, actual)
				}
			}

		})
	}
}
