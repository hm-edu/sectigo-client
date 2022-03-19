package sectigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Get(t *testing.T) {
	resp, err := GetWithoutJSONResponse(nil, nil, "")
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	resp, err = GetWithoutJSONResponse(nil, NewClient(nil, "", "", ""), "")
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func Test_Post(t *testing.T) {
	value := make(chan int)
	resp, err := PostWithoutJSONResponse(context.Background(), NewClient(http.DefaultClient, "", "", ""), "", value)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}
