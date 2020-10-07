package sms77api

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

const (
	VinTelekom = "+4915126716517"
)

var client = New(os.Getenv("SMS77_API_KEY"))

func TestNew(t *testing.T) {
	name := reflect.TypeOf(Sms77Api{
		apiKey: "",
	}).Name()

	if name != "Sms77API" {
		t.Errorf("Unexpected struct, got %s wanted %s", name, "Sms77API")
	}
}

func AssertIsPositive(propertyName string, number interface{}, t *testing.T) bool {
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
		t.Errorf("%s should be positive, but got %f", propertyName, number)
	}

	return invalid
}

func AssertIsLengthy(propertyName string, string string, t *testing.T) bool {
	if len(string) == 0 {
		t.Errorf("string %s should not be empty", propertyName)

		return false
	}

	return true
}

func TestSms77API_Analytics(t *testing.T) {
	res, err := client.Analytics(&AnalyticsParams{})

	if err != nil {
		t.Errorf("Analytics() should not return an error, but %s", err)
	}

	for _, analytics := range res {
		AssertIsPositive("UsageEur", analytics.UsageEur, t)
		AssertIsPositive("Hlr", analytics.Hlr, t)
		AssertIsPositive("Direct", analytics.Direct, t)
		AssertIsPositive("Economy", analytics.Economy, t)
		AssertIsPositive("Inbound", analytics.Inbound, t)
		AssertIsPositive("Mnp", analytics.Mnp, t)
		AssertIsPositive("Voice", analytics.Voice, t)

		AssertIsLengthy("Date", *analytics.Date, t)
	}
}

func TestSms77API_Balance(t *testing.T) {
	res, err := client.Balance()

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

	res, err := client.Contacts(ContactsParams{Action: "read"})

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

func TestSms77API_Lookup(t *testing.T) {
	lookup := func(typ string, json bool) interface{} {
		params := LookupParams{
			Type:   typ,
			Number: VinTelekom,
		}
		if json {
			params.Json = true
		}

		res, err := client.Lookup(params)

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
	assert := func(format string) {
		r, e := client.Pricing(PricingParams{Country: "de", Format: format})

		if e != nil {
			t.Errorf("Pricing() should not return an error, but %s", e)
		}

		switch res := r.(type) {
		case string:
			AssertIsLengthy("res", res, t)
		case PricingResponse:
			AssertIsPositive("CountCountries", res.CountCountries, t)
			AssertIsPositive("CountNetworks", res.CountNetworks, t)

			for _, country := range res.Countries {
				AssertIsLengthy("Country[n].CountryCode", country.CountryCode, t)
				AssertIsLengthy("Country[n].CountryName", country.CountryName, t)
				AssertIsLengthy("Country[n].CountryPrefix", country.CountryPrefix, t)

				for _, network := range country.Networks {
					AssertIsLengthy("Country.Network[n].NetworkName", network.NetworkName, t)
					AssertIsPositive("Country.Network[n].Price", network.Price, t)
				}
			}
		default:
			t.Errorf("Pricing() should return JSON or CSV, but received nil")
		}
	}

	assert("json")

	assert("csv")
}

func TestSms77API_Sms(t *testing.T) {
	sms := func(json bool) interface{} {
		params := SmsParams{
			To:   VinTelekom,
			Text: "Hey friend",
			From: "Go-Test",
		}

		if json {
			params.Json = true
		}

		res, err := client.Sms(params)

		if err != nil {
			t.Errorf("Sms() should not return an error, but %s", err)
		}

		if res == nil {
			t.Errorf("Sms() should not be nil")
		}

		return res
	}

	res := sms(true).(SmsResponse)

	AssertIsLengthy("Success", res.Success, t)
	AssertIsLengthy("Debug", res.Debug, t)
	AssertIsLengthy("SmsType", res.SmsType, t)
	AssertIsPositive("Balance", res.Balance, t)
	if len(res.Messages) == 0 {
		t.Errorf("Messages should not be empty")
	}
	AssertIsPositive("TotalPrice", res.TotalPrice, t)
}

func TestSms77API_Status(t *testing.T) {
	res, err := client.Status(StatusParams{MessageId: 77130164658})

	if err != nil {
		t.Errorf("Status() should not return an error, but %s", err)
	}

	if res == nil {
		t.Errorf("Status() should not return nil")
	}

	lines := strings.Split(*res, "\n")
	AssertIsLengthy("CODE", lines[0], t)
	AssertIsLengthy("DATETIME", lines[1], t)
}

func TestSms77API_ValidateForVoice(t *testing.T) {
	res, err := client.ValidateForVoice(ValidateForVoiceParams{Number: VinTelekom})

	if err != nil {
		t.Errorf("ValidateForVoice() should not return an error, but %s", err)
	}

	_, err = strconv.Atoi(res.Code)
	if err != nil {
		t.Errorf("Code should be numeric, but %s", err)
	}
}

func TestSms77API_Voice(t *testing.T) {
	voice := func(xml bool) interface{} {
		params := VoiceParams{To: VinTelekom, Text: "Hey friend", From: "Go-Test"}
		if xml {
			params.Xml = true
		}

		res, err := client.Voice(params)

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
