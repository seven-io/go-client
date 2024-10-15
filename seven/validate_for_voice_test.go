package seven

import (
	a "github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func validateForVoice(res *ValidateForVoiceResponse, err error, t *testing.T) {
	if nil == err {
		if testIsDummy {
			a.True(t, res.Success)
		} else {
			_, err = strconv.Atoi(res.Code)
			a.NotNil(t, err)
		}
	} else {
		a.Nil(t, res)
	}
}

func TestValidateForVoiceResource_Get(t *testing.T) {
	r, e := client.ValidateForVoice.Get(ValidateForVoiceParams{Number: "491716992343", Callback: ""})

	validateForVoice(r, e, t)
}

func TestValidateForVoiceResource_Get_Fail(t *testing.T) {
	r, e := client.ValidateForVoice.Get(ValidateForVoiceParams{Number: ""})

	validateForVoice(r, e, t)
}
