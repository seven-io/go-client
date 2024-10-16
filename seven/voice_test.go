package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
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
	} else {
		a.Nil(t, v)
	}
}
