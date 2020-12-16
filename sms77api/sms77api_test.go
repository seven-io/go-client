package sms77api

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	const expected = "Sms77API"

	name := reflect.TypeOf(client).Name()

	if name != expected {
		t.Errorf("Unexpected struct, got %s wanted %s", name, expected)
	}

	AssertIsLengthy("apiKey", client.apiKey, t)
}