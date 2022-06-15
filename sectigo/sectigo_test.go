package sectigo

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_Get(t *testing.T) {
	// nolint:staticcheck
	resp, err := GetWithoutJSONResponse(nil, nil, "")
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	logger, _ := zap.NewProduction()
	// nolint:staticcheck
	resp, err = GetWithoutJSONResponse(nil, NewClient(nil, logger, "", "", ""), "")
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}

func Test_Post(t *testing.T) {
	value := make(chan int)
	logger, _ := zap.NewProduction()
	resp, err := PostWithoutJSONResponse(context.Background(), NewClient(http.DefaultClient, logger, "", "", ""), "", value)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
}
