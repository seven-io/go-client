package sms77api

import (
	"strconv"
	"testing"
)

func TestSms77API_ValidateForVoice(t *testing.T) {
	res, err := client.ValidateForVoice.Get(ValidateForVoiceParams{Number: VinTelekom})

	if err != nil {
		t.Errorf("ValidateForVoice() should not return an error, but %s", err)
	}

	if dummy {
		AssertIsTrue("success", res.Success, t)
	} else {
		_, err = strconv.Atoi(res.Code)
		if err != nil {
			t.Errorf("Code should be numeric, but %s", err)
		}
	}
}
