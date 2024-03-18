package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	a.IsType(t, client, &API{})
	a.Exactly(t, testOptions, client.Options)
}
