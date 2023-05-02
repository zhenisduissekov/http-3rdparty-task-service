package validate

import (
	"testing"

	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
)

func Test_validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		desc string
		args service.AssignTaskReq
		want ErrorResponse
	}{
		{
			name: "test1",
			desc: "validate method len too short",
			args: service.AssignTaskReq{
				Method: "GE",
				Url:    "http://google.com",
			},
			want: ErrorResponse{
				Tag:           "min",
				Value:         "3",
				ReceivedValue: "GE",
				FailedField:   "AssignTaskReq.Method",
			},
		},
		{
			name: "test2",
			desc: "validate method url missing",
			args: service.AssignTaskReq{
				Method: "GET",
			},
			want: ErrorResponse{
				Tag:           "required",
				Value:         "",
				ReceivedValue: "",
				FailedField:   "AssignTaskReq.Url",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.args)
			if !sliceContains(err, tt.want) {
				for _, e := range err {
					t.Errorf("received = %v, but want %v", e, tt.want)
				}
			}
		})
	}

}

func sliceContains(slice []*ErrorResponse, item ErrorResponse) bool {
	for _, s := range slice {
		if *s == item {
			return true
		}
	}
	return false
}
