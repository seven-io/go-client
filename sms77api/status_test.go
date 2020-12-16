package sms77api

import (
	"strings"
	"testing"
)

func TestSms77API_Status(t *testing.T) {
	assert := func(messageId int64) []string {
		status, err := client.Status.Post(StatusParams{MessageId: messageId})
		var lines []string

		if nil == err {
			lines = strings.Split(*status, "\n")
		} else {
			AssertEquals("status", status, nil, t)
		}

		return lines
	}

	lines := assert(77131931120)
	AssertIsLengthy("CODE", lines[0], t)
	AssertIsLengthy("DATETIME", lines[1], t)

	lines = assert(0)
	AssertEquals("API_CODE", lines[0], "901", t)
}
