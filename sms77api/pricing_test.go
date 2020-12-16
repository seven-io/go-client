package sms77api

import (
	"fmt"
	"testing"
)

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
