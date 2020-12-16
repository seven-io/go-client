package sms77api

import "testing"

func TestSms77API_Voice(t *testing.T) {
	voice := func(xml bool) interface{} {
		params := VoiceParams{To: VinTelekom, Text: "Hey friend", From: "Go-Test"}
		if xml {
			params.Xml = true
		}

		res, err := client.Voice.Post(params)

		if err != nil {
			t.Errorf("Voice() should not return an error, but %s", err)
		}

		if res == nil {
			t.Errorf("Voice() should return a string, but received nil")
		}

		AssertIsLengthy("response", *res, t)

		return res
	}

	voice(false)
}
