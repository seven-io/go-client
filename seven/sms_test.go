package seven

import (
	"fmt"
	a "github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

var testSmsBaseParams = SmsBaseParams{
	Delay:               timestampString(),
	Files:               nil,
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
	testJson(t, testSmsBaseParams, true)
}

func TestSmsResource_Json_Files(t *testing.T) {
	var params = testSmsBaseParams
	params.Text = ""
	params.Files = []SmsFile{}
	var linePrefix = "TestFile"
	var contents = stringToBase64("HI2U")
	var password = "locked"

	for start := 0; start < 1; start++ {
		var validity = uint8(start + 1)
		var fileName = fmt.Sprintf("test%d.txt", start)
		params.Text += fmt.Sprintf("%s%d [[%s]]\n", linePrefix, start, fileName)
		params.Files = append(params.Files, SmsFile{
			Contents: contents,
			Name:     fileName,
			Password: &password,
			Validity: &validity,
		})
	}
	params.Text = strings.TrimSuffix(params.Text, "\n")

	var json = testJson(t, params, false)
	var msgLines = strings.Split(json.Messages[0].Text, "\n")
	a.Equal(t, len(params.Files), len(msgLines))

	for i, line := range msgLines {
		var cols = strings.Split(line, " ")
		a.Equal(t, 2, len(cols))

		a.Equal(t, fmt.Sprintf("%s%d", linePrefix, i), cols[0])

		var url = parseURL(cols[1])
		a.Equal(t, "https", url.Scheme)
		a.Equal(t, "ul.gl", url.Host)
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

func testJson(t *testing.T, params SmsBaseParams, assertText bool) *SmsResponse {
	json, err := client.Sms.Json(params)

	if nil == err {
		var debug = testIsDummy
		var Debug, _ = strconv.ParseBool(json.Debug)

		a.NotNil(t, json.Balance)
		a.Equal(t, debug, Debug)
		a.Equal(t, "direct", json.SmsType)
		a.Equal(t, StatusCodeSuccess, json.Success)
		a.Equal(t, len(strings.Split(params.To, ",")), len(json.Messages))
		for _, msg := range json.Messages {
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
			if assertText {
				a.Equal(t, params.Text, msg.Text)
			}

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
				a.Equal(t, int64(1234567890), id)
			} else {
				a.GreaterOrEqual(t, id, int64(1))
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

			var expensed, _ = toFloat(index, 1)
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
			}
		}
	} else {
		a.Nil(t, text)
	}
}
