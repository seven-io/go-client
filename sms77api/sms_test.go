package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

var testSmsBaseParams = SmsBaseParams{
	Debug:               true,
	Delay:               strconv.FormatInt(time.Now().Unix(), 10),
	Flash:               true,
	ForeignId:           "GoTestForeignId",
	From:                "Go-Test",
	Label:               "GoTestLabel",
	NoReload:            false,
	PerformanceTracking: true,
	Text:                "Hey friend",
	To:                  VinTelekom,
	Ttl:                 320000,
	Udh:                 "GoTestUserDataHeader",
	Unicode:             false,
	Utf8:                false,
}

func TestSmsResource_Json(t *testing.T) {
	json, err := client.Sms.Json(testSmsBaseParams)

	if nil == err {
		var debug = testSmsBaseParams.Debug
		if testIsDummy {
			debug = true
		}
		var Debug, _ = strconv.ParseBool(json.Debug)

		a.NotNil(t, json.Balance)
		a.Equal(t, debug, Debug)
		a.Equal(t, "direct", json.SmsType)
		a.Equal(t, StatusCodeSuccess, json.Success)
		a.NotEmpty(t, json.Messages)
		for _, msg := range json.Messages {
			a.Equal(t, "gsm", msg.Encoding)
			a.Nil(t, msg.Error)
			a.Nil(t, msg.ErrorText)
			a.Equal(t, int64(1), msg.Parts)
			a.Equal(t, testSmsBaseParams.To, msg.Recipient)
			a.Equal(t, testSmsBaseParams.From, msg.Sender)
			a.True(t, msg.Success)
			if nil != msg.Messages {
				for _, msg2 := range *msg.Messages {
					a.NotEmpty(t, msg2)
				}
			}
			a.Equal(t, testSmsBaseParams.Text, msg.Text)

			if debug {
				a.Nil(t, msg.Id)
				a.Equal(t, float64(0), msg.Price)
			} else {
				var id, _ = strconv.ParseInt(*msg.Id, 10, 0)
				a.GreaterOrEqual(t, 1, id)
				a.Greater(t, 0, msg.Price)
			}
		}
		a.GreaterOrEqual(t, json.TotalPrice, float64(0))
	} else {
		a.Equal(t, &SmsResponse{}, json)
	}
}

func TestSmsResource_Text(t *testing.T) {
	testText(t, SmsTextParams{
		SmsBaseParams: testSmsBaseParams,
	})
}

func TestSmsResource_Text_Detailed(t *testing.T) {
	testText(t, SmsTextParams{
		Details:       true,
		SmsBaseParams: testSmsBaseParams,
	})
}

func TestSmsResource_Text_With_Id(t *testing.T) {
	testText(t, SmsTextParams{
		ReturnMessageId: true,
		SmsBaseParams:   testSmsBaseParams,
	})
}

func TestSmsResource_Text_With_Id_Detailed(t *testing.T) {
	testText(t, SmsTextParams{
		Details:         true,
		ReturnMessageId: true,
		SmsBaseParams:   testSmsBaseParams,
	})
}

func parseFloat(line string) (float64, error) {
	return strconv.ParseFloat(line, 10)
}

func testText(t *testing.T, params SmsTextParams) {
	text, err := client.Sms.Text(params)

	if nil == err {
		var index = 0
		var lines = strings.Split(*text, "\n")
		var code = lines[index]
		index++

		a.Equal(t, StatusCodeSuccess, StatusCode(code))

		if params.ReturnMessageId {
			var id, _ = strconv.ParseInt(lines[index], 10, 0)

			if testIsDummy {
				a.Equal(t, 1234567890, id)
			} else {
				a.GreaterOrEqual(t, id, 1)
			}

			index++
		}

		if params.Details {
			parseLine := func(rowIndex int, lineIndex int) string {
				return strings.Split(lines[rowIndex], " ")[lineIndex]
			}

			toBool := func(rowIndex int, lineIndex int) (bool, error) {
				return strconv.ParseBool(parseLine(rowIndex, lineIndex))
			}

			toFloat := func(rowIndex int, lineIndex int) (float64, error) {
				return parseFloat(parseLine(rowIndex, lineIndex))
			}

			var expensed, _ = parseFloat(lines[index])
			index++
			var price, _ = toFloat(index, 1)
			index++
			var _, balanceError = toFloat(index, 1)
			index++
			var text = strings.TrimLeft(lines[index], "Text: ")
			index++
			var typ = parseLine(index, 1)
			index++
			var flash, _ = toBool(index, 2)
			index++
			var encoding = parseLine(index, 1)
			index++
			var gsm0338, _ = toBool(index, 1)
			index++
			var debug, _ = toBool(index, 1)

			a.Equal(t, index+1, len(lines))
			a.Equal(t, true, gsm0338)
			a.Equal(t, "gsm", encoding)
			a.Equal(t, params.Flash, flash)
			a.Equal(t, "direct", typ)
			a.Equal(t, params.Text, text)
			a.GreaterOrEqual(t, price, float64(0))
			a.Nil(t, balanceError)

			if testIsDummy {
				a.Zero(t, expensed)
				a.Equal(t, true, debug)
			} else {
				a.NotZero(t, expensed)
				a.Equal(t, params.Debug, debug)
			}
		}
	} else {
		a.Nil(t, text)
	}
}
