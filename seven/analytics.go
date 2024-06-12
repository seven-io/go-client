package seven

import (
	"context"
	"encoding/json"
)

type AnalyticsParams struct {
	End         string `json:"end,omitempty"`
	Label       string `json:"label,omitempty"`
	Start       string `json:"start,omitempty"`
	Subaccounts string `json:"subaccounts,omitempty"`
}

type Analytics struct {
	Hlr      int     `json:"hlr"`
	Inbound  int     `json:"inbound"`
	Mnp      int     `json:"mnp"`
	Rcs      int     `json:"rcs"`
	Sms      int     `json:"sms"`
	Voice    int     `json:"voice"`
	UsageEur float64 `json:"usage_eur"`
}

type AnalyticsByDate struct {
	Analytics
	Date *string `json:"date"`
}

type AnalyticsByLabel struct {
	Analytics
	Label string `json:"label"`
}

type AnalyticsBySubaccount struct {
	Analytics
	Account string `json:"account"`
}

type AnalyticsByCountry struct {
	Analytics
	Country string `json:"country"`
}

type AnalyticsResource resource

func (api *AnalyticsResource) ByCountry(p *AnalyticsParams) (o []AnalyticsByCountry, err error) {
	return api.ByCountryContext(context.Background(), p)
}

func (api *AnalyticsResource) ByCountryContext(ctx context.Context, p *AnalyticsParams) (o []AnalyticsByCountry, err error) {
	res, err := api.client.request(ctx, "analytics/country", "GET", p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	if nil != err {
		return nil, err
	}

	return o, nil
}

func (api *AnalyticsResource) ByDate(p *AnalyticsParams) (o []AnalyticsByDate, err error) {
	return api.ByDateContext(context.Background(), p)
}

func (api *AnalyticsResource) ByDateContext(ctx context.Context, p *AnalyticsParams) (o []AnalyticsByDate, err error) {
	res, err := api.client.request(ctx, "analytics/date", "GET", p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	if nil != err {
		return nil, err
	}

	return o, nil
}

func (api *AnalyticsResource) ByLabel(p *AnalyticsParams) (o []AnalyticsByLabel, err error) {
	return api.ByLabelContext(context.Background(), p)
}

func (api *AnalyticsResource) ByLabelContext(ctx context.Context, p *AnalyticsParams) (o []AnalyticsByLabel, err error) {
	res, err := api.client.request(ctx, "analytics/label", "GET", p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	if nil != err {
		return nil, err
	}

	return o, nil
}

func (api *AnalyticsResource) BySubaccount(p *AnalyticsParams) (o []AnalyticsBySubaccount, err error) {
	return api.BySubaccountContext(context.Background(), p)
}

func (api *AnalyticsResource) BySubaccountContext(ctx context.Context, p *AnalyticsParams) (o []AnalyticsBySubaccount, err error) {
	res, err := api.client.request(ctx, "analytics/subaccount", "GET", p)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	if nil != err {
		return nil, err
	}

	return o, nil
}
