package seven

import (
	"encoding/json"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestVoiceResource_Dispatch(t *testing.T) {
	v, e := client.Voice.Dispatch(VoiceParams{
		To:   "491716992343",
		Text: `<Response><Play digits="1wwwwww4"></Play><Say>Hello Sir!</Say></Response>`,
		From: "491771783130",
	})

	if nil == e {
		var msg = v.Messages[0]

		if testIsDummy {
			a.Equal(t, "100", v.Success)
			exceptedId := int64(123456789)
			a.Equal(t, &exceptedId, msg.Id)
			a.Equal(t, 0.0, v.TotalPrice)
		} else {
			a.NotEmpty(t, v.Success)
			a.NotEmpty(t, msg.Id)
			a.NotEmpty(t, v.TotalPrice)
		}

		client.Voice.Hangup(VoiceHangupParams{CallIdentifier: *msg.Id})
	} else {
		a.Nil(t, v)
	}
}

func TestVoiceMessage_UnmarshalJSON(t *testing.T) {
	expectedNumber := int64(1384013)
	tests := []struct {
		name     string
		data     []byte
		excepted *int64
	}{
		{`id="1384013"`, []byte(`{"id": "1384013"}`), &expectedNumber},
		{`id=1384013`, []byte(`{"id": 1384013}`), &expectedNumber},
		{`id=nil`, []byte(`{"id": null}`), nil},
		{`id=missing`, []byte(`{}`), nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			msg := &VoiceMessage{}
			err := json.Unmarshal(test.data, &msg)
			if a.NoError(t, err) {
				a.Equal(t, test.excepted, msg.Id)
			}
		})
	}
}
