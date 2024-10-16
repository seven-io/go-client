package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func lookup(typ string, json bool, t *testing.T) interface{} {
	res, err := client.Lookup.Get(LookupParams{
		Json:   json,
		Number: "491716992343",
		Type:   typ,
	})

	if err == nil {
		a.NotNil(t, res)
	}

	if res == nil {
		a.Nil(t, err)
	}

	return res
}

func TestLookupFormat(t *testing.T) {
	res, err := client.Lookup.Format(LookupParams{
		Number: "491716992343",
	})

	if err == nil {
		a.NotNil(t, res)
	}
	if res == nil {
		a.Nil(t, err)
	}

	a.NotEmpty(t, res.National)
	a.NotEmpty(t, res.Carrier)
	a.NotEmpty(t, res.CountryCode)
	a.NotEmpty(t, res.CountryName)
	a.NotEmpty(t, res.International)
	a.NotEmpty(t, res.InternationalFormatted)
	a.NotEmpty(t, res.NetworkType)
}

func TestLookupCnam(t *testing.T) {
	res, err := client.Lookup.Cnam(LookupParams{
		Number: "491716992343",
	})

	if err == nil {
		a.NotNil(t, res)
	}
	if res == nil {
		a.Nil(t, err)
	}

	a.NotEmpty(t, res.Code)
	a.NotEmpty(t, res.Name)
	a.NotEmpty(t, res.Number)
	a.NotEmpty(t, res.Success)
}

func TestLookupHlr(t *testing.T) {
	hlr := lookup("hlr", false, t).(*LookupHlrResponse)
	a.NotEmpty(t, hlr.CountryCode)
	a.NotEmpty(t, hlr.CountryName)
	a.NotEmpty(t, hlr.CountryPrefix)
	a.NotEmpty(t, hlr.CurrentCarrier.Country)
	a.NotEmpty(t, hlr.CurrentCarrier.Name)
	a.NotEmpty(t, hlr.CurrentCarrier.NetworkCode)
	a.NotEmpty(t, hlr.CurrentCarrier.NetworkType)
	a.NotEmpty(t, hlr.InternationalFormatNumber)
	a.NotEmpty(t, hlr.InternationalFormatted)
	a.NotEmpty(t, hlr.LookupOutcomeMessage)
	a.NotEmpty(t, hlr.NationalFormatNumber)
	a.NotEmpty(t, hlr.OriginalCarrier.Country)
	a.NotEmpty(t, hlr.OriginalCarrier.Name)
	a.NotEmpty(t, hlr.OriginalCarrier.NetworkCode)
	a.NotEmpty(t, hlr.OriginalCarrier.NetworkType)
	a.NotEmpty(t, hlr.Ported)
	a.NotEmpty(t, hlr.Reachable)
	a.NotEmpty(t, hlr.Roaming)
	a.NotEmpty(t, hlr.StatusMessage)
	a.NotEmpty(t, hlr.ValidNumber)

	if nil != hlr.CountryCodeIso3 {
		a.NotEmpty(t, *hlr.CountryCodeIso3)
	}
}

func TestLookupMnp(t *testing.T) {
	a.NotEmpty(t, lookup("mnp", false, t).(string))
}

func TestLookupMnpJson(t *testing.T) {
	mnp := lookup("mnp", true, t).(*LookupMnpResponse)
	a.Greater(t, mnp.Code, int64(0))
	a.NotEmpty(t, mnp.Mnp.Country, 0)
	a.NotEmpty(t, mnp.Mnp.InternationalFormatted, 0)
	a.NotEmpty(t, mnp.Mnp.Mccmnc, 0)
	a.NotEmpty(t, mnp.Mnp.NationalFormat, 0)
	a.NotEmpty(t, mnp.Mnp.Network, 0)
	a.NotEmpty(t, mnp.Mnp.Number, 0)
	a.Greater(t, mnp.Price, float64(0))
}
