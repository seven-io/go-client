package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

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
	res, err := client.Lookup.Hlr(LookupParams{
		Number: "491716992343",
	})

	if err == nil {
		a.NotNil(t, res)
	}
	if res == nil {
		a.Nil(t, err)
	}

	a.NotEmpty(t, res.CountryCode)
	a.NotEmpty(t, res.CountryName)
	a.NotEmpty(t, res.CountryPrefix)
	a.NotEmpty(t, res.CurrentCarrier.Country)
	a.NotEmpty(t, res.CurrentCarrier.Name)
	a.NotEmpty(t, res.CurrentCarrier.NetworkCode)
	a.NotEmpty(t, res.CurrentCarrier.NetworkType)
	a.NotEmpty(t, res.InternationalFormatNumber)
	a.NotEmpty(t, res.InternationalFormatted)
	a.NotEmpty(t, res.LookupOutcomeMessage)
	a.NotEmpty(t, res.NationalFormatNumber)
	a.NotEmpty(t, res.OriginalCarrier.Country)
	a.NotEmpty(t, res.OriginalCarrier.Name)
	a.NotEmpty(t, res.OriginalCarrier.NetworkCode)
	a.NotEmpty(t, res.OriginalCarrier.NetworkType)
	a.NotEmpty(t, res.Ported)
	a.NotEmpty(t, res.Reachable)
	a.NotEmpty(t, res.Roaming)
	a.NotEmpty(t, res.StatusMessage)
	a.NotEmpty(t, res.ValidNumber)

	if nil != res.CountryCodeIso3 {
		a.NotEmpty(t, *res.CountryCodeIso3)
	}
}

func TestLookupMnp(t *testing.T) {
	res, err := client.Lookup.Mnp(LookupParams{
		Number: "491716992343",
	})

	if err == nil {
		a.NotNil(t, res)
	}
	if res == nil {
		a.Nil(t, err)
	}

	a.Greater(t, res.Code, int64(0))
	a.NotEmpty(t, res.Mnp.Country, 0)
	a.NotEmpty(t, res.Mnp.InternationalFormatted, 0)
	a.NotEmpty(t, res.Mnp.Mccmnc, 0)
	a.NotEmpty(t, res.Mnp.NationalFormat, 0)
	a.NotEmpty(t, res.Mnp.Network, 0)
	a.NotEmpty(t, res.Mnp.Number, 0)
	a.Greater(t, res.Price, float64(0))
}
