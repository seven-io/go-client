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
		k, v := pickMapByKey(&json.Success, StatusCodes)
		a.NotNil(t, k)
		a.NotNil(t, v)
		a.NotEmpty(t, json.Debug)
		a.NotEmpty(t, json.SmsType)
		a.GreaterOrEqual(t, json.Balance, float64(0))
		a.NotEmpty(t, json.Messages)
		a.GreaterOrEqual(t, json.TotalPrice, float64(0))
	} else {
		a.Equal(t, &SmsResponse{}, json)
	}
}

func TestSmsResource_Text(t *testing.T) {
	code, err := client.Sms.Text(SmsTextParams{
		SmsBaseParams:   testSmsBaseParams,
	})
	if nil == err {
		a.Equal(t, "100", *code)
	} else {
		a.Nil(t, code)
	}
}

func TestSmsResource_Text_Detailed(t *testing.T) {
	var params = SmsTextParams{
		Details:         true,
		SmsBaseParams:   testSmsBaseParams,
	}
	text, err := client.Sms.Text(params)
	if nil == err {
		var lines = strings.Split(*text,"\n")
		var code = lines[0]
		var expensed, _ = strconv.ParseFloat(lines[1], 10)
		var price, _ = strconv.ParseFloat(strings.Split(lines[2]," ")[1], 10)
		var _, balanceError = strconv.ParseFloat(strings.Split(lines[3]," ")[1], 10)
		var text = strings.TrimLeft(lines[4], "Text: ")
		var typ =  strings.Split(lines[5]," ")[1]
		var flash, _ = strconv.ParseBool(strings.Split(lines[6]," ")[2])
		var encoding =  strings.Split(lines[7]," ")[1]
		var gsm0338, _ = strconv.ParseBool(strings.Split(lines[8]," ")[1])
		var debug, _ = strconv.ParseBool(strings.Split(lines[9]," ")[1])

		a.Equal(t, 10, len(lines))
		a.Equal(t, "100", code)
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
	} else {
		a.Nil(t, text)
	}
}