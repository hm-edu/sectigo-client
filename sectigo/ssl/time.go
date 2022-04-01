package ssl

import (
	"strings"
	"time"
)

// JSONDate is a wrapper around the time struct with a customized implementation of the json.Unmarshaler interface.
type JSONDate struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in the format MM/DD/YYYY.
func (t *JSONDate) UnmarshalJSON(buf []byte) error {
	val := strings.Trim(string(buf), `"`)
	tt, err := time.Parse("01/02/2006", val)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}
