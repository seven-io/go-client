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

type SmsParams struct {
	SmsBaseParams
	SmsTextParams
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
	SmsType    string  `json:"sms_type"`
	Success    string  `json:"success"`
	TotalPrice float64 `json:"total_price"`
}

type SmsResource resource

func (api *SmsResource) Request(p interface{}) (*string, error) {
	res, err := api.client.request("sms", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *SmsResource) Text(p SmsTextParams) (*string, error) {
	res, err := api.Request(p)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (api *SmsResource) Json(p SmsBaseParams) (*SmsResponse, error) {
	type SmsJsonParams struct {
		SmsBaseParams
		Json bool `json:"json,omitempty"`
	}

	res, err := api.Request(SmsJsonParams{
		SmsBaseParams: p,
		Json:          true,
	})

	if err != nil {
		return nil, err
	}

	var js = &SmsResponse{}

	if err := json.Unmarshal([]byte(*res), js); err != nil {
		return js, err
	}

	return js, nil
}
