package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func lookup(typ string, json bool, t *testing.T) interface{} {
	res, err := client.Lookup.Post(LookupParams{
		Json:   json,
		Number: VinTelekom,
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
	format := lookup("format", false, t).(*LookupFormatResponse)
	a.NotEmpty(t, format.National)
	a.NotEmpty(t, format.Carrier)
	a.NotEmpty(t, format.CountryCode)
	a.NotEmpty(t, format.CountryName)
	a.NotEmpty(t, format.International)
	a.NotEmpty(t, format.InternationalFormatted)
	a.NotEmpty(t, format.NetworkType)
}

func TestLookupCnam(t *testing.T) {
	cnam := lookup("cnam", false, t).(*LookupCnamResponse)
	a.NotEmpty(t, cnam.Code)
	a.NotEmpty(t, cnam.Name)
	a.NotEmpty(t, cnam.Number)
	a.NotEmpty(t, cnam.Success)
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
