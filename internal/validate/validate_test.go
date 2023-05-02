package validate

import (
	"fmt"
"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate(t *testing.T) {
	sample := struct {
		FailedField   string `validate:"required" min="7" max="10"`
		Tag           string `validate:"required" len="5"`
		Value         string `validate:"required" uppercase"`
		ReceivedValue string `validate:"required" numeric"`
	}{
		FailedField:   "test",
		Tag:           "test",
		Value:         "test",
		ReceivedValue: "test",
	}

	err := Validate(sample)
	want := ErrorResponse{
		FailedField: "test",
	}
	assert.Equal(t, err, want)

}


func Test_validate_MethoTooShort(t *testing.T) {
	items := struct {
    	Method  string            `json:"method" validate:"required,min=3,max=6,alphanum,uppercase" example:"GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD, CONNECT, TRACE"`
    	Url     string            `json:"url" validate:"required" example:"http://google.com"`
    	Headers map[string]string `json:"headers" validate:"omitempty" example:"\"Authentication\": \"Basic bG9naW46cGFzc3dvcmQ=\""`
    	ReqBody []byte            `json:"body" validate:"omitempty" example:"{\"name\":\"John\"}"`
    	Status  string            `json:"status" validate:"omitempty" example:"done/in_process/error/new"`
    }{
		Method: "GE",
		Url: "http://google.com",
    }
	
	err := Validate(items)
	want := ErrorResponse{
		FailedField: "Method",
	}
	assert.Equal(t, want, err)
	for i, v := rangeerr {
		fmt.Println("here", i, v)
	}
}