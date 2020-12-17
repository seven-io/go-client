package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"strconv"
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
	csv, err := client.Sms.Text(SmsTextParams{
		Details:         true,
		ReturnMessageId: true,
		SmsBaseParams:   testSmsBaseParams,
	})
	if nil == err {
		a.NotEmpty(t, *csv)
	} else {
		a.Nil(t, csv)
	}
}
