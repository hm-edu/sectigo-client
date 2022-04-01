package sectigo

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func stringFromResponse(err error, resp *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	return &bodyString, nil
}

func formatParams[T any](q *T) (string, error) {
	params := ""
	if q != nil {
		values, err := query.Values(q)
		if err != nil {
			return "", err
		}
		params = values.Encode()
		if len(params) != 0 {
			params = fmt.Sprintf("?%v", params)
		}
	}
	return params, nil
}
