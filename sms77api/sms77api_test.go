package sms77api

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	VinTelekom = "+4915126716517"
)

var client, dummy = GetClient()

func GetClient() (*Sms77API, bool) {
	var dummy = true
	var apiKey = os.Getenv("SMS77_DUMMY_API_KEY")

	if "" == apiKey {
		apiKey = os.Getenv("SMS77_API_KEY")

		dummy = false
	}

	return New(apiKey), dummy
}

func AssertIsPositive(descriptor string, number interface{}, t *testing.T) bool {
	invalid := false

	switch number.(type) {
	case int:
		invalid = number.(int) < 0
	case float32:
		invalid = number.(float32) < 0
	case float64:
		invalid = number.(float64) < 0
	}

	if invalid {
		t.Errorf("%s should be positive, but got %f", descriptor, number)
	}

	return invalid
}

func AssertIsTrue(descriptor string, value interface{}, t *testing.T) bool {
	if true != value {
		t.Errorf("%s should be true, but is not", descriptor)
	}

	return true
}

func AssertIsNil(descriptor string, value interface{}, t *testing.T) bool {
	if nil != value {
		t.Errorf("%s should be nil, but is not", descriptor)
	}

	return true
}

func AssertIsLengthy(descriptor string, value string, t *testing.T) bool {
	if len(value) == 0 {
		t.Errorf("string %s should not be empty", descriptor)

		return false
	}

	return true
}

func AssertEquals(descriptor string, actual interface{}, expected interface{}, t *testing.T) bool {
	if expected != actual {
		t.Errorf("%s should match %v but received %v", descriptor, expected, actual)

		return false
	}

	return true
}

func AssertInArray(descriptor string, needle interface{}, haystack interface{}, t *testing.T) bool {
	slice := reflect.ValueOf(haystack)
	c := slice.Len()

	for i := 0; i < c; i++ {
		if needle == slice.Index(i).Interface() {
			return true
		}
	}

	t.Errorf("%s with value %s should be included in %s", descriptor, needle, slice)

	return false
}

func TestNew(t *testing.T) {
	const expected = "Sms77API"

	name := reflect.TypeOf(client).Name()

	if name != expected {
		t.Errorf("Unexpected struct, got %s wanted %s", name, expected)
	}

	AssertIsLengthy("apiKey", client.apiKey, t)
}

func TestSms77API_Analytics(t *testing.T) {
	res, err := client.Analytics.Get(&AnalyticsParams{})

	if err != nil {
		t.Errorf("Analytics() should not return an error, but %s", err)
	}

	for _, analytics := range res {
		AssertIsLengthy("Date", *analytics.Date, t)
		AssertIsPositive("Direct", analytics.Direct, t)
		AssertIsPositive("Economy", analytics.Economy, t)
		AssertIsPositive("Hlr", analytics.Hlr, t)
		AssertIsPositive("Inbound", analytics.Inbound, t)
		AssertIsPositive("Mnp", analytics.Mnp, t)
		AssertIsPositive("UsageEur", analytics.UsageEur, t)
		AssertIsPositive("Voice", analytics.Voice, t)
	}
}

func TestSms77API_Balance(t *testing.T) {
	res, err := client.Balance.Get()

	if err != nil {
		t.Errorf("Balance() should not return an error, but %s", err)
	}

	if res == nil {
		t.Errorf("Balance() should return a float64 value, but received nil")
	}

	AssertIsPositive("Balance()", res, t)
}

func TestSms77API_Contacts(t *testing.T) {
	assertContact := func(c Contact) bool {
		invalid := false

		if c.Id < 0 {
			t.Errorf("Every Contact must have a positive ID")

			invalid = true
		}

		return invalid
	}

	toStruct := func(c string) Contact {
		c = strings.TrimSpace(c)
		arr := strings.Split(c, ";")

		id, err := strconv.ParseInt(strings.ReplaceAll(arr[0], "\"", ""), 10, 64)
		if err != nil {
			t.Errorf("Contacts.Id should should be a int64 value, but %s", err)
		}

		return Contact{
			Id:    id,
			Nick:  strings.ReplaceAll(arr[1], "\"", ""),
			Phone: strings.ReplaceAll(arr[2], "\"", ""),
		}
	}

	res, err := client.Contacts.ReadCsv(ContactsReadParams{})

	if err != nil {
		t.Errorf("Contacts() should not return an error, but %s", err)
	}

	if res == nil {
		t.Errorf("Contacts() should return a string value, but received nil")
	}

	for _, csvContact := range strings.Split(strings.TrimSpace(*res), "\n") {
		assertContact(toStruct(csvContact))
	}
}

func TestSms77API_Hooks(t *testing.T) {
	request := func(params HooksParams) interface{} {
		res, err := client.Hooks.Request(params)

		log.Print(res)

		if err != nil {
			t.Errorf("Hooks() should not return an error, but %s", err)
		}

		if res == nil {
			t.Errorf("Hooks() should return json, but received nil")
		}

		return res
	}

	hooks := request(HooksParams{Action: HooksActionRead}).(*HooksReadResponse)

	if hooks.Success && hooks.Hooks != nil {
		for _, hook := range hooks.Hooks {
			AssertIsLengthy("Created", hook.Created, t)
			AssertIsLengthy("Id", hook.Id, t)
			AssertIsLengthy("TargetUrl", hook.TargetUrl, t)
			AssertInArray("EventType", hook.EventType,
				[...]HookEventType{HookEventTypeSmsStatus, HookEventTypeVoiceStatus, HookEventTypeInboundSms}, t)
			AssertInArray("RequestMethod", hook.RequestMethod,
				[...]HookRequestMethod{HookRequestMethodGet, HookRequestMethodPost}, t)
		}
	}

	subscribed := request(HooksParams{
		Action:        HooksActionSubscribe,
		EventType:     HookEventTypeInboundSms,
		RequestMethod: HookRequestMethodGet,
		TargetUrl:     fmt.Sprintf("https://test.tld/go-client/%d", time.Now().Unix()),
	}).(*HooksSubscribeResponse)

	AssertIsPositive("Id", subscribed.Id, t)

	if true == subscribed.Success {
		subscribed := request(HooksParams{
			Action: HooksActionUnsubscribe,
			Id:     subscribed.Id,
		}).(*HooksUnsubscribeResponse)

		AssertIsTrue("Success", subscribed.Success, t)
	}
}

func TestSms77API_Lookup(t *testing.T) {
	lookup := func(typ string, json bool) interface{} {
		params := LookupParams{
			Type:   typ,
			Number: VinTelekom,
		}
		if json {
			params.Json = true
		}

		res, err := client.Lookup.Post(params)

		if err != nil {
			t.Errorf("Lookup() should not return an error, but %s", err)
		}

		if res == nil {
			t.Errorf("Lookup() should return a string or json, but received nil")
		}

		return res
	}

	format := lookup("format", false).(*LookupFormatResponse)
	AssertIsLengthy("National", format.National, t)
	AssertIsLengthy("Carrier", format.Carrier, t)
	AssertIsLengthy("CountryCode", format.CountryCode, t)
	AssertIsLengthy("CountryName", format.CountryName, t)
	AssertIsLengthy("International", format.International, t)
	AssertIsLengthy("InternationalFormatted", format.InternationalFormatted, t)
	AssertIsLengthy("NetworkType", format.NetworkType, t)

	cnam := lookup("cnam", false).(*LookupCnamResponse)
	AssertIsLengthy("Code", cnam.Code, t)
	AssertIsLengthy("Name", cnam.Name, t)
	AssertIsLengthy("Number", cnam.Number, t)
	AssertIsLengthy("Success", cnam.Success, t)

	hlr := lookup("hlr", false).(*LookupHlrResponse)
	AssertIsLengthy("CountryCode", hlr.CountryCode, t)
	AssertIsLengthy("CountryName", hlr.CountryName, t)
	AssertIsLengthy("CountryPrefix", hlr.CountryPrefix, t)
	AssertIsLengthy("CurrentCarrier.Country", hlr.CurrentCarrier.Country, t)
	AssertIsLengthy("CurrentCarrier.Name", hlr.CurrentCarrier.Name, t)
	AssertIsLengthy("CurrentCarrier.NetworkCode", hlr.CurrentCarrier.NetworkCode, t)
	AssertIsLengthy("CurrentCarrier.NetworkType", hlr.CurrentCarrier.NetworkType, t)
	AssertIsLengthy("InternationalFormatNumber", hlr.InternationalFormatNumber, t)
	AssertIsLengthy("InternationalFormatted", hlr.InternationalFormatted, t)
	AssertIsLengthy("LookupOutcomeMessage", hlr.LookupOutcomeMessage, t)
	AssertIsLengthy("NationalFormatNumber", hlr.NationalFormatNumber, t)
	AssertIsLengthy("OriginalCarrier.Country", hlr.OriginalCarrier.Country, t)
	AssertIsLengthy("OriginalCarrier.Name", hlr.OriginalCarrier.Name, t)
	AssertIsLengthy("OriginalCarrier.NetworkCode", hlr.OriginalCarrier.NetworkCode, t)
	AssertIsLengthy("OriginalCarrier.NetworkType", hlr.OriginalCarrier.NetworkType, t)
	AssertIsLengthy("Ported", hlr.Ported, t)
	AssertIsLengthy("Reachable", hlr.Reachable, t)
	AssertIsLengthy("Roaming", hlr.Roaming, t)
	AssertIsLengthy("StatusMessage", hlr.StatusMessage, t)
	AssertIsLengthy("ValidNumber", hlr.ValidNumber, t)

	if hlr.CountryCodeIso3 != nil {
		AssertIsLengthy("CountryCodeIso3", *hlr.CountryCodeIso3, t)
	}

	AssertIsLengthy("response", lookup("mnp", false).(string), t)

	mnp := lookup("mnp", true).(*LookupMnpResponse)
	AssertIsPositive("Code", mnp.Code, t)
	AssertIsLengthy("Mnp.Country", mnp.Mnp.Country, t)
	AssertIsLengthy("Mnp.InternationalFormatted", mnp.Mnp.InternationalFormatted, t)
	AssertIsLengthy("Mnp.Mccmnc", mnp.Mnp.Mccmnc, t)
	AssertIsLengthy("Mnp.NationalFormat", mnp.Mnp.NationalFormat, t)
	AssertIsLengthy("Mnp.Network", mnp.Mnp.Network, t)
	AssertIsLengthy("Mnp.Number", mnp.Mnp.Number, t)
	AssertIsPositive("Price", mnp.Price, t)
}

func TestSms77API_Pricing(t *testing.T) {
	var pricingParams = PricingParams{Country: "de"}

	var json, jsonError = client.Pricing.Json(pricingParams)
	if nil == jsonError {
		AssertIsNil("jsonError", jsonError, t)
		AssertIsPositive("CountCountries", json.CountCountries, t)
		AssertIsPositive("CountNetworks", json.CountNetworks, t)

		for n, country := range json.Countries {
			AssertIsLengthy(fmt.Sprintf("Country[%d].CountryCode", n), country.CountryCode, t)
			AssertIsLengthy(fmt.Sprintf("Country[%d].CountryName", n), country.CountryName, t)
			AssertIsLengthy(fmt.Sprintf("Country[%d].CountryPrefix", n), country.CountryPrefix, t)

			for nn, network := range country.Networks {
				AssertIsLengthy(fmt.Sprintf("Country[%d].Network[%d].NetworkName", n, nn), network.NetworkName, t)
				AssertIsPositive(fmt.Sprintf("Country[%d].Network[%d].Price", n, nn), network.Price, t)
			}
		}
	} else {
		AssertEquals("res", json, "", t)
	}

	var csv, csvErr = client.Pricing.Csv(pricingParams)
	if nil == csvErr {
		AssertIsLengthy("res", csv, t)
	} else {
		AssertEquals("res", csv, "", t)
	}
}

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

func TestSms77API_Status(t *testing.T) {
	assert := func(messageId int64) []string {
		status, err := client.Status.Post(StatusParams{MessageId: messageId})
		var lines []string

		if nil == err {
			lines = strings.Split(*status, "\n")
		} else {
			AssertEquals("status", status, nil, t)
		}

		return lines
	}

	lines := assert(77131931120)
	AssertIsLengthy("CODE", lines[0], t)
	AssertIsLengthy("DATETIME", lines[1], t)

	lines = assert(0)
	AssertEquals("API_CODE", lines[0], "901", t)
}

func TestSms77API_ValidateForVoice(t *testing.T) {
	res, err := client.ValidateForVoice.Get(ValidateForVoiceParams{Number: VinTelekom})

	if err != nil {
		t.Errorf("ValidateForVoice() should not return an error, but %s", err)
	}

	if dummy {
		AssertIsTrue("success", res.Success, t)
	} else {
		_, err = strconv.Atoi(res.Code)
		if err != nil {
			t.Errorf("Code should be numeric, but %s", err)
		}
	}
}

func TestSms77API_Voice(t *testing.T) {
	voice := func(xml bool) interface{} {
		params := VoiceParams{To: VinTelekom, Text: "Hey friend", From: "Go-Test"}
		if xml {
			params.Xml = true
		}

		res, err := client.Voice.Post(params)

		if err != nil {
			t.Errorf("Voice() should not return an error, but %s", err)
		}

		if res == nil {
			t.Errorf("Voice() should return a string, but received nil")
		}

		AssertIsLengthy("response", *res, t)

		return res
	}

	voice(false)
}
