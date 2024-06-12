package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestBalance(t *testing.T) {
	res, err := client.Balance.Get()

	if err != nil {
		t.Errorf("Balance() should not return an error, but %s", err)
	}

	if res == nil {
		t.Errorf("Balance() should return a float64 value, but received nil")
	}

	//a.GreaterOrEqual(t, res.Amount, 0.0)
	a.NotEmpty(t, res.Currency)
}

func TestBalanceBad(t *testing.T) {
	r, e := testBadClient.Balance.Get()
	AssertKeylessCall(t, r, e)
}
