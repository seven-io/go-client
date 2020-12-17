package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func journal(journals interface{}, err error, t *testing.T) interface{} {
	if nil != err {
		a.Nil(t, journals)
	}

	return journals
}

func TestJournalInbound(t *testing.T) {
	r, e := client.Journal.Inbound(&JournalParams{})

	journal(r, e, t)
}

func TestJournalOutbound(t *testing.T) {
	r, e := client.Journal.Outbound(&JournalParams{})

	for _, j := range journal(r, e, t).([]JournalOutbound) {
		a.Greater(t, len(j.Connection), 0)
		a.Greater(t, len(j.Type), 0)
	}
}

func TestJournalReplies(t *testing.T) {
	r, e := client.Journal.Replies(&JournalParams{})

	journal(r, e, t)
}

func TestJournalVoice(t *testing.T) {
	r, e := client.Journal.Voice(&JournalParams{})

	for _, j := range journal(r, e, t).([]JournalVoice) {
		a.Greater(t, len(j.Duration), 0)
		a.Greater(t, len(j.Status), 0)
	}
}
