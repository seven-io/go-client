package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func testVoiceBuildParams(p VoiceParams) VoiceParams {
	text := "Just testing ;-)"

	if p.Xml {
		text = `<Response><Play digits="1wwwwww4"></Play><Say>Hello Sir!</Say></Response>`
	}

	p.From = VinEplus
	p.Text = text
	p.To = VinTelekom

	return p
}

func testVoiceJson(p VoiceParams, t *testing.T) {
	v, e := client.Voice.Json(testVoiceBuildParams(p))

	if nil == e {
		testVoiceAssert(p, v, t)
	} else {
		a.Nil(t, v)
	}
}

func testVoiceText(p VoiceParams, t *testing.T) {
	res, err := client.Voice.Text(testVoiceBuildParams(p))

	if nil == err {
		testVoiceAssert(p, makeVoice(*res), t)
	} else {
		a.Nil(t, res)
	}
}

func testVoiceAssert(p VoiceParams, v Voice, t *testing.T) {
	if testIsDummy || p.Debug {
		var x = Voice{Code: 100, Cost: 0, Id: 123456789}

		a.Equal(t, x.Code, v.Code)
		a.Equal(t, x.Id, v.Id)
		a.Equal(t, x.Cost, v.Cost)
	} else {
		a.NotEmpty(t, v.Code)
		a.NotEmpty(t, v.Id)
		a.NotEmpty(t, v.Cost)
	}
}

func TestVoiceResource_Text(t *testing.T) {
	testVoiceText(VoiceParams{}, t)
}

func TestVoiceResource_Text_Xml(t *testing.T) {
	testVoiceText(VoiceParams{Xml: true}, t)
}

func TestVoiceResource_Json(t *testing.T) {
	testVoiceJson(VoiceParams{}, t)
}

func TestVoiceResource_Json_Xml(t *testing.T) {
	testVoiceJson(VoiceParams{Xml: true}, t)
}
