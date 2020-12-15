package sms77api

type VoiceParams struct {
	To   string `json:"to"`
	Text string `json:"text"`
	Xml  bool   `json:"xml,omitempty"`
	From string `json:"from,omitempty"`
}

type VoiceResource resource

func (api *VoiceResource) Post(p VoiceParams) (*string, error) {
	res, err := api.client.request("voice", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
