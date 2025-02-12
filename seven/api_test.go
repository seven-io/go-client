package seven

import (
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a.IsType(t, client, &API{})
	a.Exactly(t, testOptions, client.Options)
}

func TestStatusText(t *testing.T) {
	text := StatusText(StatusCodeSuccess)
	a.Equal(t, StatusCodes[StatusCodeSuccess], text)

	unknownCode := StatusCode("000")
	text = StatusText(unknownCode)
	a.Contains(t, text, string(unknownCode))
}
