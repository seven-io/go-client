package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	a.IsType(t, client, &Sms77API{})
	a.Exactly(t, testOptions, client.Options)
}
