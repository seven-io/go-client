package sms77api

import "encoding/json"

type AnalyticsParams struct {
	Start       string `json:"start,omitempty"`
	End         string `json:"end,omitempty"`
	GroupBy     string `json:"group_by,omitempty"`
	Label       string `json:"label,omitempty"`
	Subaccounts string `json:"subaccounts,omitempty"`
}

type Analytics struct {
	Account *string `json:"account"`
	Country *string `json:"country"`
	Date    *string `json:"date"`
	Label   *string `json:"label"`

	Direct   int     `json:"direct"`
	Economy  int     `json:"economy"`
	Hlr      int     `json:"hlr"`
	Inbound  int     `json:"inbound"`
	Mnp      int     `json:"mnp"`
	Voice    int     `json:"voice"`
	UsageEur float64 `json:"usage_eur"`
}

type AnalyticsResponse []Analytics

type AnalyticsResource resource

func (api *AnalyticsResource) Get(p *AnalyticsParams) (AnalyticsResponse, error) {
	res, err := api.client.request("analytics", "GET", p)
	if err != nil {
		return nil, err
	}

	var js = AnalyticsResponse{}
	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}
	return js, nil
}