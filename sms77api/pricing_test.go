package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestPricingCsv(t *testing.T) {
	var params = PricingParams{Country: "DE"}
	var csv, csvErr = client.Pricing.Csv(params)
	if nil == csvErr {
		a.NotEmpty(t, csv)

		lines := strings.Split(csv, "\n")

		for i, line := range lines {
			line = strings.ReplaceAll(line, "; ", ";")
			cols := regexp.MustCompile(`"(.*?)"`).FindAllString(line, -1)
			for i, col := range cols {
				cols[i] = strings.ReplaceAll(col, `"`, "")
			}
			a.Len(t, cols, len(PricingCsvHeaders))

			if 0 == i {
				for ii, h := range PricingCsvHeaders {
					a.Equal(t, h, cols[ii])
				}
			} else {
				a.Equal(t, params.Country, cols[PricingColumnCountryCode])
				a.NotEmpty(t, cols[PricingColumnCountryName])
				a.NotEmpty(t, cols[PricingColumnCountryPrefix])
				a.NotEmpty(t, cols[PricingColumnMcc])
				a.NotEmpty(t, cols[PricingColumnNetworkName])
				a.NotEmpty(t, cols[PricingColumnPrice])
				a.NotEmpty(t, cols[PricingColumnMncs])
			}
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
