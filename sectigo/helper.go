package sectigo

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"strings"
	"time"
)

type JsonDate struct {
	time.Time
}

func (t *JsonDate) UnmarshalJSON(buf []byte) error {
	val := strings.Trim(string(buf), `"`)
	tt, err := time.Parse("2006-01-02", val)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

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
		params = fmt.Sprintf("?%v", values.Encode())
	}
	return params, nil
}
