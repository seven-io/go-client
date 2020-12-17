package sms77api

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func analytics(gB AnalyticsGroupBy, eachEntry func(Analytics), t *testing.T) {
	res, err := client.Analytics.Get(&AnalyticsParams{GroupBy: gB, Label: "all", Subaccounts: "all"})

	if nil == err {
		for _, o := range res {
			eachEntry(o)

			a.GreaterOrEqual(t, o.Direct, 0)
			a.GreaterOrEqual(t, o.Economy, 0)
			a.GreaterOrEqual(t, o.Hlr, 0)
			a.GreaterOrEqual(t, o.Inbound, 0)
			a.GreaterOrEqual(t, o.Mnp, 0)
			a.GreaterOrEqual(t, o.UsageEur, float64(0))
			a.GreaterOrEqual(t, o.Voice, 0)
		}
	} else {
		a.Nil(t, res)
	}
}

func TestAnalyticsCountry(t *testing.T) {
	analytics(AnalyticsGroupByCountry, func(o Analytics) {
		a.NotNil(t, o.Country)
		a.Nil(t, o.Account)
		a.Nil(t, o.Date)
		a.Nil(t, o.Label)
	}, t)
}

func TestAnalyticsDate(t *testing.T) {
	analytics(AnalyticsGroupByDate, func(o Analytics) {
		a.Greater(t, len(*o.Date), 0)
		a.Nil(t, o.Account)
		a.Nil(t, o.Country)
		a.Nil(t, o.Label)
	}, t)
}

func TestAnalyticsLabel(t *testing.T) {
	analytics(AnalyticsGroupByLabel, func(o Analytics) {
		a.NotNil(t, o.Label)
		a.Nil(t, o.Account)
		a.Nil(t, o.Country)
		a.Nil(t, o.Date)
	}, t)
}

func TestAnalyticsSubaccount(t *testing.T) {
	analytics(AnalyticsGroupBySubaccount, func(o Analytics) {
		a.Greater(t, len(*o.Account), 0)
		a.Nil(t, o.Country)
		a.Nil(t, o.Date)
		a.Nil(t, o.Label)
	}, t)
}
