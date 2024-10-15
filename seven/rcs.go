package seven

import (
	"context"
	"encoding/json"
	"strconv"
)

type RcsFallback string

const (
	RcsFallbackSms     RcsFallback = "sms"
	RcsFallbackWebview RcsFallback = "webview"
)

type RcsEvent string

const (
	RcsEventIsTyping RcsEvent = "IS_TYPING"
	RcsEventRead     RcsEvent = "READ"
)

type RcsEventParams struct {
	To        string   `json:"to,omitempty"`
	MessageId string   `json:"msg_id,omitempty"`
	Event     RcsEvent `json:"event"`
}

type RcsParams struct {
	Delay               string      `json:"delay,omitempty"`
	ForeignId           string      `json:"foreign_id,omitempty"`
	From                string      `json:"from,omitempty"`
	Label               string      `json:"label,omitempty"`
	PerformanceTracking bool        `json:"performance_tracking,omitempty"`
	Text                string      `json:"text"`
	To                  string      `json:"to"`
	TTL                 int64       `json:"ttl,omitempty"`
	RcsFallback         RcsFallback `json:"fallback,omitempty"`
}

type RcsResource resource

type RcsResponse struct {
	Debug      string       `json:"debug"`
	Balance    float64      `json:"balance"`
	Messages   []RcsMessage `json:"messages"`
	SmsType    string       `json:"sms_type"`
	Success    StatusCode   `json:"success"`
	TotalPrice float64      `json:"total_price"`
}

type RcsMessage struct {
	Channel   string    `json:"channel"`
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

type RcsEventResponse struct {
	Success bool `json:"success"`
}

type RcsDeletionResponse struct {
	Success bool `json:"success"`
}

func (api *RcsResource) Dispatch(p RcsParams) (o *RcsResponse, err error) {
	return api.DispatchContext(context.Background(), p)
}

func (api *RcsResource) DispatchContext(ctx context.Context, p RcsParams) (o *RcsResponse, err error) {
	res, err := api.client.request(ctx, "rcs/messages", "POST", p)

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	return
}

func (api *RcsResource) Delete(id uint) (o *RcsDeletionResponse, err error) {
	return api.DeleteContext(context.Background(), id)
}

func (api *RcsResource) DeleteContext(ctx context.Context, id uint) (o *RcsDeletionResponse, err error) {
	res, err := api.client.request(ctx, "rcs/messages/"+strconv.Itoa(int(id)), "DELETE", nil)

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	return
}

func (api *RcsResource) Event(params RcsEventParams) (o *RcsEventResponse, err error) {
	return api.EventContext(context.Background(), params)
}

func (api *RcsResource) EventContext(ctx context.Context, params RcsEventParams) (o *RcsEventResponse, err error) {
	res, err := api.client.request(ctx, "rcs/events", string(HttpMethodPost), params)

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal([]byte(res), &o)

	return
}
