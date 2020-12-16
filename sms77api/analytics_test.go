package sms77api

import "testing"

func TestSms77API_Analytics(t *testing.T) {
	res, err := client.Analytics.Get(&AnalyticsParams{})

	if err != nil {
		t.Errorf("Analytics() should not return an error, but %s", err)
	}

	for _, analytics := range res {
		AssertIsLengthy("Date", *analytics.Date, t)
		AssertIsPositive("Direct", analytics.Direct, t)
		AssertIsPositive("Economy", analytics.Economy, t)
		AssertIsPositive("Hlr", analytics.Hlr, t)
		AssertIsPositive("Inbound", analytics.Inbound, t)
		AssertIsPositive("Mnp", analytics.Mnp, t)
		AssertIsPositive("UsageEur", analytics.UsageEur, t)
		AssertIsPositive("Voice", analytics.Voice, t)
	}
}
