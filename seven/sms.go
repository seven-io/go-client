package seven

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

type SmsParams struct {
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
	Error     *string   `json:"error"`
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

func (api *SmsResource) Json(p SmsParams) (o *SmsResponse, err error) {
	return api.JsonContext(context.Background(), p)
}
func (api *SmsResource) JsonContext(ctx context.Context, p SmsParams) (o *SmsResponse, err error) {
	res, err := api.request(ctx, p)

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(*res), &o)

	return
}
