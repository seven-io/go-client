package seven

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

type VoiceHangupParams struct {
	CallIdentifier int64
}

type VoiceHangup struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
}

type Voice struct {
	Balance    float64        `json:"balance"`
	Debug      bool           `json:"debug"`
	TotalPrice float64        `json:"total_price"`
	Success    StatusCode     `json:"success"`
	Messages   []VoiceMessage `json:"messages"`
}

type VoiceMessage struct {
	Error     *StatusCode `json:"error"`
	ErrorText *string `json:"error_text"`
	Id        *int64  `json:"id"`
	Price     float64 `json:"price"`
	Recipient string  `json:"recipient"`
	Sender    string  `json:"sender"`
	Success   bool    `json:"success"`
	Text      string  `json:"text"`
}

// UnmarshalJSON is a workaround as a result of https://github.com/seven-io/go-client/issues/10.
// The Id can be a string or a number and is converted into a number here.
func (m *VoiceMessage) UnmarshalJSON(b []byte) error {
	data := make(map[string]interface{}, 0)
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	switch idVal := data["id"].(type) {
	case string:
		if id, err := strconv.ParseInt(idVal, 10, 64); err != nil {
			return err
		} else {
			data["id"] = id
		}
	}

	b, errM := json.Marshal(data)
	if errM != nil {
		return errM
	}

	// This is a trick that prevents the object from referencing itself during encoding,
	// which would result in an endless recursion.
	type messageCopy VoiceMessage
	var result messageCopy
	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}
	*m = VoiceMessage(result)

	return nil
}

type VoiceParams struct {
	To       string `json:"to"`
	Text     string `json:"text"`
	From     string `json:"from,omitempty"`
	Ringtime uint8  `json:"ringtime,omitempty"`
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

func (api *VoiceResource) Hangup(p VoiceHangupParams) (o *VoiceHangup, e error) {
	return api.HangupContext(context.Background(), p)
}

func (api *VoiceResource) HangupContext(ctx context.Context, p VoiceHangupParams) (*VoiceHangup, error) {
	endpoint := fmt.Sprintf("voice/%d/hangup", p.CallIdentifier)
	res, err := api.client.request(ctx, endpoint, "POST", nil)

	if err != nil {
		return nil, err
	}

	var js = &VoiceHangup{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return js, nil
}
