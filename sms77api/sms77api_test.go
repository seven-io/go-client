package sms77api

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	const expected = "Sms77API"

	name := reflect.TypeOf(*client).Name()

	if name != expected {
		t.Errorf("Unexpected struct, got %s wanted %s", name, expected)
	}

	AssertEquals("ApiKey", client.ApiKey, options.ApiKey, t)
	AssertEquals("Debug", client.Debug, options.Debug, t)
	AssertEquals("SentWith", client.SentWith, options.SentWith, t)
}

func TestNewFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for missing ApiKey")
		}
	}()

	New(Options{})
}
