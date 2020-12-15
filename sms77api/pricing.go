package sms77api

import "encoding/json"

type PricingResource resource

type CountryNetwork struct {
	Comment     string   `json:"comment,omitempty"`
	Features    []string `json:"features,omitempty"`
	Mcc         string   `json:"mcc,omitempty"`
	Mncs        []string `json:"mncs,omitempty"`
	NetworkName string   `json:"networkName,omitempty"`
	Price       float64  `json:"price,omitempty"`
}

type CountryPricing struct {
	CountryCode   string           `json:"countryCode,omitempty"`
	CountryName   string           `json:"countryName,omitempty"`
	CountryPrefix string           `json:"countryPrefix,omitempty"`
	Networks      []CountryNetwork `json:"networks,omitempty"`
}

type PricingFormat string

const (
	PricingFormatCsv   PricingFormat = "csv"
	PricingFormatJson  PricingFormat = "json"
)

type PricingParams struct {
	Country string `json:"country,omitempty"`
	//Format  PricingFormat `json:"format,omitempty"`
}

type PricingResponse struct {
	CountCountries int64            `json:"countCountries"`
	CountNetworks  int64            `json:"countNetworks"`
	Countries      []CountryPricing `json:"countries"`
}

type PricingApiParams struct {
	PricingParams
	Format  PricingFormat `json:"format,omitempty"`
}

const ENDPOINT = "pricing"

func (api *PricingResource) Csv(p PricingParams) (string, error) {
	res, err := api.client.request(ENDPOINT, "GET", PricingApiParams{
		PricingParams: p,
		Format: PricingFormatCsv,
	})

	if err != nil {
		return "", err
	}

	return res, nil
}

func (api *PricingResource) Json(p PricingParams) (*PricingResponse, error) {
	res, err := api.client.request(ENDPOINT, "GET", PricingApiParams{PricingParams: p})

	if err != nil {
		return nil, err
	}

	var js = &PricingResponse{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return js, nil
}

/*func (api *PricingResource) Get(p *PricingParams) (interface{}, error) {
	if "csv" == p.Format {
		return api.Csv(p)
	}

	return api.Json(p)
}*/
