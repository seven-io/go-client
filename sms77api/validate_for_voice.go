package sms77api

import "encoding/json"

type ValidateForVoiceParams struct {
	Number   string `json:"number"`
	Networks string `json:"networks,omitempty"`
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

func (api *ValidateForVoiceResource) Get(p ValidateForVoiceParams) (*ValidateForVoiceResponse, error) {
	res, err := api.client.request("validate_for_voice", "GET", p)

	if err != nil {
		return nil, err
	}

	var js = &ValidateForVoiceResponse{}

	if err := json.Unmarshal([]byte(res), js); err != nil {
		return nil, err
	}

	return js, nil
}
