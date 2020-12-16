package sms77api

import (
	"testing"
)

func TestSms77API_Journal(t *testing.T) {
	request := func(journals interface{}, err error) interface{} {
		if nil != err {
			AssertEquals("journals", journals, nil, t)
		}

		return journals
	}

	request(client.Journal.Inbound(&JournalParams{}))

	for _, journal := range request(client.Journal.Outbound(&JournalParams{})).([]JournalOutbound) {
		AssertIsLengthy("Connection", journal.Connection, t)
		AssertIsLengthy("Type", journal.Type, t)
	}

	request(client.Journal.Replies(&JournalParams{}))

	for _, journal := range request(client.Journal.Voice(&JournalParams{})).([]JournalVoice) {
		AssertIsLengthy("Duration", journal.Duration, t)
		AssertIsLengthy("Status", journal.Status, t)
	}
}
