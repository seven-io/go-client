package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func status(messageId uint64, t *testing.T) []string {
	status, err := client.Status.Post(StatusParams{MessageId: messageId})
	var lines []string

	if nil == err {
		lines = strings.Split(*status, "\n")
	} else {
		a.Nil(t, status)
	}

	return lines
}

func TestStatusResource_Post(t *testing.T) {
	journals, _ := client.Journal.Outbound(&JournalParams{})
	var id string

	if 0 == len(journals) {
		sms, _ := client.Sms.Json(SmsBaseParams{To: VinTelekom, Text: "HI"})
		id = sms.Messages[0].Id
	} else {
		id = journals[0].Id
	}

	lines := status(toUint(id, 64), t)
	a.Len(t, lines, 2)
	a.NotEmpty(t, lines[0])
	a.NotEmpty(t, lines[1])
}

func TestStatusResource_Post_Fail(t *testing.T) {
	lines := status(0, t)
	a.Len(t, lines, 1)
	a.Equal(t, StatusApiCodeInvalidMessageId, lines[0])
}
