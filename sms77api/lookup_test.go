package sms77api

import (
	"testing"
)

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
