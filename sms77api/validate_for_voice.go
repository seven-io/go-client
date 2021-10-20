package sms77api

import (
	"context"
	"encoding/json"
)

type ValidateForVoiceParams struct {
	Callback string `json:"callback"`
	Number   string `json:"number"`
}

type ValidateForVoiceResponse struct {
	Code            string  `json:"code"`
	Error           *string `json:"error"`
	FormattedOutput *string `json:"formatted_output"`
	Id              *int64  `json:"id"`
	Sender          string  `json:"sender"`
	Success         bool    `json:"success"`
	Voice           bool    `json:"voice"`
}

type ValidateForVoiceResource resource

func (api *ValidateForVoiceResource) Get(p ValidateForVoiceParams) (o *ValidateForVoiceResponse, e error) {
	return api.GetContext(context.Background(), p)
}

func (api *ValidateForVoiceResource) GetContext(ctx context.Context, p ValidateForVoiceParams) (o *ValidateForVoiceResponse, e error) {
	r, e := api.client.request(ctx, "validate_for_voice", "GET", p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), o)

	return
}
