package sms77api

import (
	"context"
	"encoding/json"
)

type SmsFile struct {
	Contents string  `json:"contents"`
	Name     string  `json:"name"`
	Validity *uint8  `json:"validity,omitempty"`
	Password *string `json:"password,omitempty"`
}

type SmsBaseParams struct {
	Debug               bool      `json:"debug,omitempty"`
	Delay               string    `json:"delay,omitempty"`
	Files               []SmsFile `json:"files,omitempty"`
	Flash               bool      `json:"flash,omitempty"`
	ForeignId           string    `json:"foreign_id,omitempty"`
	From                string    `json:"from,omitempty"`
	Label               string    `json:"label,omitempty"`
	PerformanceTracking bool      `json:"performance_tracking,omitempty"`
	Text                string    `json:"text"`
	To                  string    `json:"to"`
	Ttl                 int64     `json:"ttl,omitempty"`
	Udh                 string    `json:"udh,omitempty"`
	// Deprecated: Unicode is not necessary as the gateway automatically recognizes the coding.
	Unicode bool `json:"unicode,omitempty"`
	// Deprecated: Utf8 is deprecated.
	Utf8 bool `json:"utf8,omitempty"`
	// Deprecated: NoReload is deprecated.
	NoReload bool `json:"no_reload,omitempty"`
}

type SmsTextParams struct {
	SmsBaseParams
	// Deprecated: Details is deprecated.
	Details bool `json:"details,omitempty"`
	// Deprecated: ReturnMessageId is deprecated.
	ReturnMessageId bool `json:"return_msg_id,omitempty"`
}

type SmsResource resource

type SmsResponse struct {
	Debug      string               `json:"debug"`
	Balance    float64              `json:"balance"`
	Messages   []SmsResponseMessage `json:"messages"`
	SmsType    string               `json:"sms_type"`
	Success    StatusCode           `json:"success"`
	TotalPrice float64              `json:"total_price"`
}

type SmsResponseMessage struct {
	Encoding  string    `json:"encoding"`
	Error     *int64    `json:"error"`
	ErrorText *string   `json:"error_text"`
	Id        *string   `json:"id"`
	Label     *string   `json:"label"`
	Messages  *[]string `json:"messages,omitempty"`
	Parts     int64     `json:"parts"`
	Price     float64   `json:"price"`
	Recipient string    `json:"recipient"`
	Sender    string    `json:"sender"`
	Success   bool      `json:"success"`
	Text      string    `json:"text"`
}

func (api *SmsResource) request(ctx context.Context, p interface{}) (*string, error) {
	res, err := api.client.request(ctx, "sms", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *SmsResource) Text(p SmsTextParams) (res *string, err error) {
	return api.TextContext(context.Background(), p)
}

func (api *SmsResource) TextContext(ctx context.Context, p SmsTextParams) (res *string, err error) {
	return api.request(ctx, p)
}

func (api *SmsResource) Json(p SmsBaseParams) (o *SmsResponse, err error) {
	return api.JsonContext(context.Background(), p)
}
func (api *SmsResource) JsonContext(ctx context.Context, p SmsBaseParams) (o *SmsResponse, err error) {
	type SmsJsonParams struct {
		SmsBaseParams
		Json bool `json:"json,omitempty"`
	}

	res, err := api.request(ctx, SmsJsonParams{
		SmsBaseParams: p,
		Json:          true,
	})

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(*res), &o)

	return
}
