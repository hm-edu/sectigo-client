package sectigo

import (
	"encoding/json"
	"fmt"
	"github.com/hm-edu/sectigo-client/sectigo/misc"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatParams(t *testing.T) {
	i := 42
	val, err := formatParams(&i)
	assert.Equal(t, "", val)
	assert.NotNil(t, err)
}

type errorReader struct {
}

func (e *errorReader) Read(_ []byte) (n int, err error) {
	return 0, fmt.Errorf("Cannot read")
}

func Test_StringFromResponse(t *testing.T) {
	val, err := stringFromResponse(fmt.Errorf("Fatal"), nil)
	assert.NotNil(t, err)
	assert.Nil(t, val)

	reader := io.NopCloser(&errorReader{})
	val, err = stringFromResponse(nil, &http.Response{Body: reader})
	assert.NotNil(t, err)
	assert.Nil(t, val)
}

func TestJsonDate_UnmarshalJSON(t *testing.T) {
	d := misc.JSONDate{}

	err := json.NewDecoder(strings.NewReader("1.1.2022")).Decode(&d)
	assert.NotNil(t, err)
}
