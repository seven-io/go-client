package seven

import (
	"context"
	"encoding/json"
	"fmt"
)

type PaymentInterval string

const (
	PaymentIntervalAnnually PaymentInterval = "annually"
	PaymentIntervalMonthly  PaymentInterval = "monthly"
)

type NumbersResource resource

type Number struct {
	FriendlyName      string                  `json:"friendly_name"`
	Created           string                  `json:"created"`
	Expires           string                  `json:"expires"`
	Country           string                  `json:"country"`
	Number            string                  `json:"number"`
	Features          NumberFeatures          `json:"features"`
	Billing           NumberBilling           `json:"billing"`
	ForwardInboundSms NumberForwardInboundSms `json:"forward_sms_mo"`
}
type AvailableNumber struct {
	Country      string              `json:"country"`
	Number       string              `json:"number"`
	NumberParsed string              `json:"number_parsed"`
	Fees         AvailableNumberFees `json:"fees"`
	Features     NumberFeatures      `json:"features"`
}
type AvailableNumberFees struct {
	Monthly      AvailableNumberFeesMonthly  `json:"monthly"`
	Annually     AvailableNumberFeesAnnually `json:"annually"`
	InboundSms   float64                     `json:"sms_mo"`
	InboundVoice float64                     `json:"voice_mo"`
}
type AvailableNumberFeesMonthly struct {
	BasicCharge float64 `json:"basic_charge"`
	Setup       float64 `json:"setup"`
}
type AvailableNumberFeesAnnually struct {
	BasicCharge float64 `json:"basic_charge"`
	Setup       float64 `json:"setup"`
}
type NumberFeatures struct {
	SMS                    bool `json:"sms"`
	ApplicationToPersonSms bool `json:"a2p_sms"`
	Voice                  bool `json:"voice"`
}
type NumbersDeleteParams = struct {
	DeleteImmediately bool `json:"delete_immediately,omitempty"`
}
type NumberBilling struct {
	PaymentInterval string            `json:"payment_interval"`
	Fees            NumberBillingFees `json:"fees"`
}
type NumberBillingFees struct {
	Setup        float64 `json:"setup"`
	BasicCharge  float64 `json:"basic_charge"`
	InboundSms   float64 `json:"sms_mo"`
	InboundVoice float64 `json:"voice_mo"`
}
type NumberForwardInboundSms struct {
	Sms   NumberForwardInboundSmsBySms   `json:"sms"`
	Email NumberForwardInboundSmsByEmail `json:"email"`
}
type NumberForwardInboundSmsBySms struct {
	Numbers []string `json:"number"`
	Enabled bool     `json:"enabled"`
}
type NumberForwardInboundSmsByEmail struct {
	Addresses []string `json:"address"`
	Enabled   bool     `json:"enabled"`
}

type NumberOrderParams struct {
	Number          string          `json:"number"`
	PaymentInterval PaymentInterval `json:"payment_interval"`
}

type NumberUpdateParams struct {
	FriendlyName string   `json:"friendly_name"`
	SmsForward   []string `json:"sms_forward"`
	EmailForward []string `json:"email_forward"`
}

type NumbersGetParams = struct {
	Number string `json:"number"`
}

type NumbersActive = struct {
	ActiveNumbers []Number `json:"activeNumbers"`
}

type NumbersAvailable = struct {
	AvailableNumbers []AvailableNumber `json:"availableNumbers"`
}

type NumbersAvailableParams = struct {
	Country                        string `json:"country,omitempty"`
	FeaturesSms                    bool   `json:"features_sms,omitempty"`
	FeaturesApplicationToPersonSms bool   `json:"features_a2p_sms,omitempty"`
	FeaturesVoice                  bool   `json:"features_voice,omitempty"`
}

type NumberDeleted = struct {
	Success bool `json:"success"`
}

func (api *NumbersResource) Get(p NumbersGetParams) (c Number, e error) {
	return api.GetContext(context.Background(), p)
}

func (api *NumbersResource) GetContext(ctx context.Context, p NumbersGetParams) (c Number, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("numbers/%s", p.Number), string(HttpMethodGet), nil)

	if nil != e {
		return
	}

	json.Unmarshal([]byte(s), &c)

	return
}

func (api *NumbersResource) AvailableNumbers(params NumbersAvailableParams) (a NumbersAvailable, e error) {
	return api.AvailableNumbersContext(context.Background(), params)
}

func (api *NumbersResource) AvailableNumbersContext(ctx context.Context, p NumbersAvailableParams) (a NumbersAvailable, e error) {
	s, e := api.client.request(ctx, "numbers/available", string(HttpMethodGet), p)

	if nil != e {
		return
	}

	json.Unmarshal([]byte(s), &a)

	return
}

func (api *NumbersResource) ActiveNumbers() (a NumbersActive, e error) {
	return api.ActiveNumbersContext(context.Background())
}

func (api *NumbersResource) ActiveNumbersContext(ctx context.Context) (a NumbersActive, e error) {
	s, e := api.client.request(ctx, "numbers/active", string(HttpMethodGet), nil)

	if nil != e {
		return
	}

	json.Unmarshal([]byte(s), &a)

	return
}

func (api *NumbersResource) Order(p NumberOrderParams) (c Number, e error) {
	return api.OrderContext(context.Background(), p)
}

func (api *NumbersResource) OrderContext(ctx context.Context, p NumberOrderParams) (c Number, e error) {
	s, e := api.client.request(ctx, "numbers/order", string(HttpMethodPost), p)

	e = json.Unmarshal([]byte(s), &c)

	return
}

func (api *NumbersResource) Delete(phone string, p NumbersDeleteParams) (o NumberDeleted, e error) {
	return api.DeleteContext(context.Background(), phone, p)
}

func (api *NumbersResource) DeleteContext(ctx context.Context, phone string, p NumbersDeleteParams) (o NumberDeleted, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("numbers/active/%s", phone), string(HttpMethodDelete), p)

	e = json.Unmarshal([]byte(s), &o)

	return
}

func (api *NumbersResource) Update(number string, p NumberUpdateParams) (n Number, e error) {
	return api.UpdateContext(context.Background(), number, p)
}

func (api *NumbersResource) UpdateContext(ctx context.Context, number string, p NumberUpdateParams) (n Number, e error) {
	r, e := api.client.request(ctx, fmt.Sprintf("numbers/active/%s", number), string(HttpMethodPatch), p)
	e = json.Unmarshal([]byte(r), &n)
	return
}
