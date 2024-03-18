package seven

import (
	"context"
	"strconv"
)

type BalanceResource resource

func (api *BalanceResource) Get() (*float64, error) {
	return api.GetContext(context.Background())
}

func (api *BalanceResource) GetContext(ctx context.Context) (*float64, error) {
	res, err := api.client.request(ctx, "balance", "GET", nil)

	if err != nil {
		return nil, err
	}

	float, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return nil, err
	}

	return &float, nil
}
