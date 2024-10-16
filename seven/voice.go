package seven

import (
	"context"
	"encoding/json"
)

type Voice struct {
	Balance    float64        `json:"balance"`
	Debug      bool           `json:"debug"`
	TotalPrice float64        `json:"total_price"`
	Success    string         `json:"success"`
	Messages   []VoiceMessage `json:"messages"`
}

type VoiceMessage struct {
	Error     *string `json:"error"`
	ErrorText *string `json:"error_text"`
	Id        *string `json:"id"`
	Price     float64 `json:"price"`
	Recipient string  `json:"recipient"`
	Sender    string  `json:"sender"`
	Success   bool    `json:"success"`
	Text      string  `json:"text"`
}

type VoiceParams struct {
	To   string `json:"to"`
	Text string `json:"text"`
	From string `json:"from,omitempty"`
}

type VoiceResource resource

func (api *VoiceResource) Dispatch(p VoiceParams) (o *Voice, e error) {
	return api.DispatchContext(context.Background(), p)
}

func (api *VoiceResource) DispatchContext(ctx context.Context, p VoiceParams) (*Voice, error) {
	res, err := api.client.request(ctx, "voice", "POST", p)

	if err != nil {
		return nil, err
	}

	var js = &Voice{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return js, nil
}
