package seven

import (
	a "github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

var testRcsParams = RcsParams{
	Delay:               timestampString(),
	ForeignId:           "GoTestForeignId",
	From:                "Go-Test",
	Label:               "GoTestLabel",
	PerformanceTracking: true,
	Text:                "Hey friend",
	To:                  VinTelekom,
	TTL:                 320000,
}

func TestRcsResource_Text(t *testing.T) {
	testRcs(t, testRcsParams)
}

func TestRcsResource_Delete(t *testing.T) {
	id := uint(123)
	json, err := client.Rcs.Delete(id)

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
