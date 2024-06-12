package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyticsCountry(t *testing.T) {
	res, err := client.Analytics.ByCountry(&AnalyticsParams{Label: "all", Subaccounts: "all"})

	if nil == err {
		for _, o := range res {
			a.NotNil(t, o.Country)
			a.GreaterOrEqual(t, o.Hlr, 0)
			a.GreaterOrEqual(t, o.Inbound, 0)
			a.GreaterOrEqual(t, o.Mnp, 0)
			a.GreaterOrEqual(t, o.Rcs, 0)
			a.GreaterOrEqual(t, o.Sms, 0)
			a.GreaterOrEqual(t, o.UsageEur, float64(0))
			a.GreaterOrEqual(t, o.Voice, 0)
		}
	} else {
		a.Nil(t, res)
	}
}

func TestAnalyticsDate(t *testing.T) {
	res, err := client.Analytics.ByDate(&AnalyticsParams{Label: "all", Subaccounts: "all"})

	if nil == err {
		for _, o := range res {
			a.Greater(t, len(*o.Date), 0)
			a.GreaterOrEqual(t, o.Hlr, 0)
			a.GreaterOrEqual(t, o.Inbound, 0)
			a.GreaterOrEqual(t, o.Mnp, 0)
			a.GreaterOrEqual(t, o.Rcs, 0)
			a.GreaterOrEqual(t, o.Sms, 0)
			a.GreaterOrEqual(t, o.UsageEur, float64(0))
			a.GreaterOrEqual(t, o.Voice, 0)
		}
	} else {
		a.Nil(t, res)
	}
}

func TestAnalyticsLabel(t *testing.T) {
	res, err := client.Analytics.ByLabel(&AnalyticsParams{Label: "all", Subaccounts: "all"})

	if nil == err {
		for _, o := range res {
			a.GreaterOrEqual(t, o.Hlr, 0)
			a.GreaterOrEqual(t, o.Inbound, 0)
			a.NotNil(t, o.Label)
			a.GreaterOrEqual(t, o.Mnp, 0)
			a.GreaterOrEqual(t, o.Rcs, 0)
			a.GreaterOrEqual(t, o.Sms, 0)
			a.GreaterOrEqual(t, o.UsageEur, float64(0))
			a.GreaterOrEqual(t, o.Voice, 0)
		}
	} else {
		a.Nil(t, res)
	}
}

func TestAnalyticsSubaccount(t *testing.T) {
	res, err := client.Analytics.BySubaccount(&AnalyticsParams{Label: "all", Subaccounts: "all"})

	if nil == err {
		for _, o := range res {
			a.Greater(t, len(o.Account), 0)
			a.GreaterOrEqual(t, o.Hlr, 0)
			a.GreaterOrEqual(t, o.Inbound, 0)
			a.GreaterOrEqual(t, o.Mnp, 0)
			a.GreaterOrEqual(t, o.Rcs, 0)
			a.GreaterOrEqual(t, o.Sms, 0)
			a.GreaterOrEqual(t, o.UsageEur, float64(0))
			a.GreaterOrEqual(t, o.Voice, 0)
		}
	} else {
		a.Nil(t, res)
	}
}
