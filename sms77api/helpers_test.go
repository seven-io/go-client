package sms77api

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

const (
	VinTelekom = "+4915126716517"
)

var client, dummy = GetClient()

func GetClient() (*Sms77API, bool) {
	var dummy = true
	var apiKey = os.Getenv("SMS77_DUMMY_API_KEY")

	if "" == apiKey {
		apiKey = os.Getenv("SMS77_API_KEY")

		dummy = false
	}

	return New(apiKey), dummy
}

func AssertIsPositive(descriptor string, number interface{}, t *testing.T) bool {
	valid := false

	switch number.(type) {
	case float32:
		valid = number.(float32) > 0
	case float64:
		valid = number.(float64) > 0
	case int:
		valid = number.(int) > 0
	case int8:
		valid = number.(int8) > 0
	case int16:
		valid = number.(int16) > 0
	case int32:
		valid = number.(int32) > 0
	case int64:
		valid = number.(int64) > 0
	case uint:
		valid = number.(uint) > 0
	case uint8:
		valid = number.(uint8) > 0
	case uint32:
		valid = number.(uint32) > 0
	case uint64:
		valid = number.(uint64) > 0
	}

	if !valid {
		t.Errorf("%s should be positive, but got %f", descriptor, number)
	}

	return valid
}

func AssertIsTrue(descriptor string, value interface{}, t *testing.T) bool {
	if true != value {
		t.Errorf("%s should be true, but is not", descriptor)

		return false
	}

	return true
}

func AssertIsNil(descriptor string, value interface{}, t *testing.T) bool {
	if nil != value {
		t.Errorf("%s should be nil, but is not", descriptor)

		return false
	}

	return true
}

func AssertIsLengthy(descriptor string, value string, t *testing.T) bool {
	if 0 == len(value) {
		t.Errorf("string %s should not be empty", descriptor)

		return false
	}

	return true
}

func AssertEquals(descriptor string, actual interface{}, expected interface{}, t *testing.T) bool {
	if expected != actual {
		t.Errorf("%s should match %v but received %v", descriptor, expected, actual)

		return false
	}

	return true
}

func AssertInArray(descriptor string, needle interface{}, haystack interface{}, t *testing.T) bool {
	 if InArray(needle, haystack) {
	 	return true
	 }

	t.Errorf("%s with value %s should be included in %v", descriptor, needle, haystack)

	return false
}

func InArray(needle interface{}, haystack interface{}) bool {
	slice := reflect.ValueOf(haystack)
	c := slice.Len()

	for i := 0; i < c; i++ {
		if needle == slice.Index(i).Interface() {
			return true
		}
	}

	return false
}

func toUint64(id string) uint64 {
	n, err := strconv.ParseUint(id, 10, 64)

	if nil == err {
		return n
	}

	return 0
}
