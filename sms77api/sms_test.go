package sms77api

import (
	"strconv"
	"testing"
	"time"
)

func TestSms77API_Sms(t *testing.T) {
	baseParams := SmsBaseParams{
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

	json, jsonError := client.Sms.Json(baseParams)
	if nil == jsonError {
		AssertIsLengthy("Success", json.Success, t)
		AssertIsLengthy("Debug", json.Debug, t)
		AssertIsLengthy("SmsType", json.SmsType, t)
		AssertIsPositive("Balance", json.Balance, t)
		if len(json.Messages) == 0 {
			t.Errorf("Messages should not be empty")
		}
		AssertIsPositive("TotalPrice", json.TotalPrice, t)
	} else {
		AssertEquals("json", json, nil, t)
	}

	csv, csvError := client.Sms.Text(SmsTextParams{
		Details:         true,
		ReturnMessageId: true,
		SmsBaseParams:   baseParams,
	})
	if nil == csvError {
		AssertIsLengthy("csv", *csv, t)
	} else {
		AssertEquals("csv", csv, nil, t)
	}
}
