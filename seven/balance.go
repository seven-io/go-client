package seven

import (
	"context"
	"encoding/json"
)

type Balance struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type BalanceResource resource

func (api *BalanceResource) Get() (b *Balance, e error) {
	return api.GetContext(context.Background())
}

func (api *BalanceResource) GetContext(ctx context.Context) (b *Balance, e error) {
	res, e := api.client.request(ctx, "balance", "GET", nil)

	if e != nil {
		return nil, e
	}

	json.Unmarshal([]byte(res), &b)

	return
}
