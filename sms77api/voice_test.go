package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

type testVoiceDummyExpectation struct {
	code int
	cost float64
	id   int
}

func voice(p VoiceParams, e testVoiceDummyExpectation, t *testing.T) {
	p.From = "Go-Test"
	p.To = VinTelekom
	res, err := client.Voice.Post(p)

	if nil == err {
		a.NotEmpty(t, *res)

		lines := strings.Split(*res, "\n")
		a.Len(t, lines, 3)

		code, _ := strconv.Atoi(lines[0])
		id, _ := strconv.Atoi(lines[1])
		cost, _ := strconv.ParseFloat(lines[2], 64)

		if testIsDummy {
			a.Equal(t, e.code, code)
			a.Equal(t, e.id, id)
			a.Equal(t, e.cost, cost)
		} else {
			a.NotEmpty(t, code)
			a.NotEmpty(t, id)
			a.NotEmpty(t, cost)
		}

	} else {
		a.Nil(t, res)
	}
}

func TestVoiceResource_Post(t *testing.T) {
	voice(VoiceParams{Text: "Just testing ;-)"}, testVoiceDummyExpectation{code: 100, cost: 0, id: 123456789}, t)
}

func TestVoiceResource_Post_Xml(t *testing.T) {
	xml := `
		<?xml version="1.0" encoding="UTF-8"?>
			<Response>
				<Say voice="woman" language="en-EN">
					Your glasses are ready for pickup.
				</Say>
			<Record maxlength="20" />
		</Response>
	`
	voice(VoiceParams{Text: xml, Xml: true}, testVoiceDummyExpectation{code: 203, cost: 0.1, id: 0}, t)
}
