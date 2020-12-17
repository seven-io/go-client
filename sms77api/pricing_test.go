package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPricingCsv(t *testing.T) {
	var csv, csvErr = client.Pricing.Csv(PricingParams{Country: "de"})
	if nil == csvErr {
		a.NotEmpty(t, csv)

		csv = strings.ReplaceAll(csv, "\"", "")
		lines := strings.Split(csv, "\n")

		for i, line := range lines {
			cols := strings.Split(line, ";")

			a.Len(t, cols, len(PricingCsvHeaders))

			for ii, h := range PricingCsvHeaders {
				if 0 == i {
					a.Equal(t, h, cols[ii])
				} else {
					a.NotEqual(t, h, cols[ii])
				}
			}

			a.NotEmpty(t, cols[PricingHeaderCountryCode])
			a.NotEmpty(t, cols[PricingHeaderCountryName])
			a.NotEmpty(t, cols[PricingHeaderCountryPrefix])
			a.NotEmpty(t, cols[PricingHeaderMcc])
			a.NotEmpty(t, cols[PricingHeaderMncs])
			a.NotEmpty(t, cols[PricingHeaderNetworkName])
			a.NotEmpty(t, cols[PricingHeaderPrice])
		}
	} else {
		a.Empty(t, csv)
	}
}

func TestPricingJson(t *testing.T) {
	var json, jsonError = client.Pricing.Json(PricingParams{Country: "fr"})
	if nil == jsonError {
		a.GreaterOrEqual(t, json.CountCountries, int64(0))
		a.GreaterOrEqual(t, json.CountNetworks, int64(0))
		a.Equal(t, int64(len(json.Countries)), json.CountCountries)

		var networkCount int64

		for _, country := range json.Countries {
			a.NotEmpty(t, country.CountryCode)
			a.NotEmpty(t, country.CountryName)
			a.NotEmpty(t, country.CountryPrefix)

			for _, network := range country.Networks {
				networkCount++
				a.NotEmpty(t, network.NetworkName)
				a.GreaterOrEqual(t, network.Price, float64(0))
			}
		}

		a.Equal(t, json.CountNetworks, networkCount)
	} else {
		a.Empty(t, json)
	}
}
