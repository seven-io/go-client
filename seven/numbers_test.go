package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestNumbers(t *testing.T) {
	availableNumbers, e := client.Numbers.AvailableNumbers(NumbersAvailableParams{
		Country:                        "DE",
		FeaturesSms:                    true,
		FeaturesApplicationToPersonSms: true,
		FeaturesVoice:                  false,
	})
	if e != nil {
		t.Errorf(e.Error())
	}
	var availableNumber = availableNumbers.AvailableNumbers[0]

	number, e := client.Numbers.Order(NumberOrderParams{
		Number:          availableNumber.Number,
		PaymentInterval: PaymentIntervalMonthly,
	})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.Equal(t, availableNumber.Number, number.Number)

	updated, e := client.Numbers.Update(number.Number, NumberUpdateParams{
		FriendlyName: "New Friendly Name",
		SmsForward:   nil,
		EmailForward: nil,
	})
	a.NotEqual(t, number.FriendlyName, updated.FriendlyName)

	actives, e := client.Numbers.ActiveNumbers()
	a.NotEmpty(t, actives.ActiveNumbers)

	single, e := client.Numbers.Get(NumbersGetParams{Number: updated.Number})
	a.NotEmpty(t, single)

	deleted, e := client.Numbers.Delete(single.Number, NumbersDeleteParams{DeleteImmediately: true})
	a.True(t, deleted.Success)
}
