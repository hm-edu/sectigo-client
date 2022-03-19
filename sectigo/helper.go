package sectigo

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// JSONDate is a wrapper around the time struct with a customized implementation of the json.Unmarshaler interface.
type JSONDate struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in the format YYYY-MM-DD.
func (t *JSONDate) UnmarshalJSON(buf []byte) error {
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
