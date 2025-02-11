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
			a.Equal(t, 0, msg.Id)
			a.Equal(t, 0, v.TotalPrice)
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

func TestVoice_UnmarshalJSON(t *testing.T) {
	// The following example is taken from the API documentation and should be able to be transformed without errors.
	// See: https://docs.seven.io/de/rest-api/endpunkte/voice
	apiExample := []byte(`
{
  "success": "100",
  "total_price": 0.045,
  "balance": 3509.236,
  "debug": false,
  "messages": [
    {
      "id": 1384013,
      "sender": "sender",
      "recipient": "49176123456789",
      "text": "Hallo Welt!",
      "price": 0.045,
      "success": true,
      "error": null,
      "error_text": null
    }
  ]
}`)

	js := &Voice{}
	err := json.Unmarshal(apiExample, js)
	if a.NoError(t, err) {
		a.Len(t, js.Messages, 1)
	}
}

func TestVoiceHangup_UnmarshalJSON(t *testing.T) {
	// The following example is taken from the API documentation and should be able to be transformed without errors.
	// See: https://docs.seven.io/de/rest-api/endpunkte/voice
	apiExample := []byte(`
{
  "success": true,
  "error": null
}`)

	js := &VoiceHangup{}
	err := json.Unmarshal(apiExample, js)
	if a.NoError(t, err) {
		a.True(t, js.Success)
		a.Nil(t, js.Error)
	}

}
