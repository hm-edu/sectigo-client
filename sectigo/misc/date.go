package misc

import (
	"strings"
	"time"
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
