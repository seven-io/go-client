package seven

import (
	a "github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestRcsResource_Event(t *testing.T) {
	params := RcsEventParams{
		Event: RcsEventIsTyping,
		To:    "+491716992343",
	}
	json, err := client.Rcs.Event(params)
	a.Nil(t, err)
	a.True(t, json.Success)
}

func TestRcsResource_Text(t *testing.T) {
	params := RcsParams{
		Delay:               "2050-12-12 00:00:00",
		ForeignId:           "GoTestForeignId",
		From:                "single",
		Label:               "GoTestLabel",
		PerformanceTracking: true,
		Text:                "Hey friend",
		To:                  "+4915229617888",
		TTL:                 320000,
	}
	testRcs(t, params)
}

func TestRcsResource_Delete(t *testing.T) {
	json, err := client.Rcs.Delete(uint(123))

	a.Equal(t, nil == err, json.Success)
}

func testRcs(t *testing.T, params RcsParams) *RcsResponse {
	json, err := client.Rcs.Dispatch(params)

	if nil == err {
		var debug = testIsDummy
		var Debug, _ = strconv.ParseBool(json.Debug)

		a.NotNil(t, json.Balance)
		a.Equal(t, debug, Debug)
		a.Equal(t, "direct", json.SmsType)
		a.Equal(t, StatusCodeSuccess, json.Success)
		a.Equal(t, len(strings.Split(params.To, ",")), len(json.Messages))
		for _, msg := range json.Messages {
			a.Equal(t, "RCS", msg.Channel)
			a.Equal(t, "gsm", msg.Encoding)
			a.Nil(t, msg.Error)
			a.Nil(t, msg.ErrorText)
			if nil != msg.Messages {
				for _, msgMsg := range *msg.Messages {
					a.NotEmpty(t, msgMsg)
				}
			}
			a.Equal(t, int64(1), msg.Parts)
			a.Equal(t, params.To, msg.Recipient)
			a.Equal(t, params.From, msg.Sender)
			a.True(t, msg.Success)
			a.Equal(t, params.Text, msg.Text)

			if debug {
				a.Nil(t, msg.Id)
				a.Equal(t, float64(0), msg.Price)
			} else {
				var id, _ = strconv.ParseInt(*msg.Id, 10, 0)
				a.GreaterOrEqual(t, id, int64(1))
				a.Greater(t, msg.Price, float64(0))
			}
		}
		a.GreaterOrEqual(t, json.TotalPrice, float64(0))
	} else {
		a.Equal(t, &SmsResponse{}, json)
	}

	return json
}
