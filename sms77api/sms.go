package sms77api

import "encoding/json"

type SmsBaseParams struct {
	Debug               bool   `json:"debug,omitempty"`
	Delay               string `json:"delay,omitempty"`
	Flash               bool   `json:"flash,omitempty"`
	ForeignId           string `json:"foreign_id,omitempty"`
	From                string `json:"from,omitempty"`
	Label               string `json:"label,omitempty"`
	NoReload            bool   `json:"no_reload,omitempty"`
	PerformanceTracking bool   `json:"performance_tracking,omitempty"`
	Text                string `json:"text"`
	To                  string `json:"to"`
	Ttl                 int64  `json:"ttl,omitempty"`
	Udh                 string `json:"udh,omitempty"`
	Unicode             bool   `json:"unicode,omitempty"`
	Utf8                bool   `json:"utf8,omitempty"`
}

type SmsTextParams struct {
	SmsBaseParams
	Details         bool `json:"details,omitempty"`
	ReturnMessageId bool `json:"return_msg_id,omitempty"`
}

type SmsResponse struct {
	Debug    string  `json:"debug"`
	Balance  float64 `json:"balance"`
	Messages []struct {
		Encoding  string    `json:"encoding"`
		Error     string    `json:"error"`
		ErrorText string    `json:"error_text"`
		Id        string    `json:"id"`
		Messages  *[]string `json:"messages,omitempty"`
		Parts     int64     `json:"parts"`
		Price     float64   `json:"price"`
		Recipient string    `json:"recipient"`
		Sender    string    `json:"sender"`
		Success   bool      `json:"success"`
		Text      string    `json:"text"`
	} `json:"messages"`
	SmsType    string     `json:"sms_type"`
	Success    StatusCode `json:"success"`
	TotalPrice float64    `json:"total_price"`
}

type SmsResource resource

func (api *SmsResource) request(p interface{}) (*string, error) {
	res, err := api.client.request("sms", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *SmsResource) Text(p SmsTextParams) (res *string, err error) {
	return api.request(p)
}

func (api *SmsResource) Json(p SmsBaseParams) (o *SmsResponse, err error) {
	type SmsJsonParams struct {
		SmsBaseParams
		json bool
	}

	res, err := api.request(SmsJsonParams{
		SmsBaseParams: p,
		json:          true,
	})

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(*res), &o)

	return
}
