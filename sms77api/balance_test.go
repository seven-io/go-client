package sms77api

import "testing"

func TestSms77API_Balance(t *testing.T) {
	res, err := client.Balance.Get()

	if err != nil {
		t.Errorf("Balance() should not return an error, but %s", err)
	}

	if res == nil {
		t.Errorf("Balance() should return a float64 value, but received nil")
	}

	AssertIsPositive("Balance()", res, t)
}