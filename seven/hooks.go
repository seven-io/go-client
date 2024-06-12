package seven

import (
	"context"
	"encoding/json"
	"fmt"
)

type HookEventType string

const (
	HookEventTypeSmsStatus   HookEventType = "dlr"
	HookEventTypeVoiceStatus HookEventType = "voice_status"
	HookEventTypeInboundSms  HookEventType = "sms_mo"
	HookEventTypeTracking    HookEventType = "tracking"
	HookEventTypeAll         HookEventType = "all"
	HookEventTypeRcs         HookEventType = "rcs"
)

type HookRequestMethod string

const (
	HookRequestMethodGet  HookRequestMethod = "GET"
	HookRequestMethodJson HookRequestMethod = "JSON"
	HookRequestMethodPost HookRequestMethod = "POST"
)

type Hook struct {
	Created       string            `json:"created"`
	Enabled       bool              `json:"enabled"`
	EventFilter   *string           `json:"event_filter"`
	EventType     HookEventType     `json:"event_type"`
	Id            string            `json:"id"`
	RequestMethod HookRequestMethod `json:"request_method"`
	TargetUrl     string            `json:"target_url"`
}

type HooksSubscribeParams struct {
	EventFilter   string            `json:"event_filter,omitempty"`
	EventType     HookEventType     `json:"event_type"`
	RequestMethod HookRequestMethod `json:"request_method,omitempty"`
	TargetUrl     string            `json:"target_url"`
}

type HooksListResponse struct {
	Hooks   []Hook `json:"hooks"`
	Success bool   `json:"success"`
}

type HooksUnsubscribeResponse struct {
	Success bool `json:"success"`
}

type HooksSubscribeResponse struct {
	Id      uint `json:"id"`
	Success bool `json:"success"`
}

type HooksResource resource

func (api *HooksResource) Subscribe(p HooksSubscribeParams) (r *HooksSubscribeResponse, e error) {
	return api.SubscribeContext(context.Background(), p)
}

func (api *HooksResource) SubscribeContext(ctx context.Context, p HooksSubscribeParams) (r *HooksSubscribeResponse, e error) {
	res, e := api.client.request(ctx, "hooks", string(HttpMethodPost), p)

	if e != nil {
		return nil, e
	}

	if e := json.Unmarshal([]byte(res), &r); e != nil {
		return nil, e
	}

	return
}

func (api *HooksResource) Unsubscribe(id uint) (r *HooksUnsubscribeResponse, e error) {
	return api.UnsubscribeContext(context.Background(), id)
}

func (api *HooksResource) UnsubscribeContext(ctx context.Context, id uint) (r *HooksUnsubscribeResponse, e error) {
	res, e := api.client.request(ctx, fmt.Sprintf("hooks?id=%d", id), string(HttpMethodDelete), nil)

	if e != nil {
		return nil, e
	}

	if e := json.Unmarshal([]byte(res), &r); e != nil {
		return nil, e
	}

	return
}

func (api *HooksResource) List() (r *HooksListResponse, e error) {
	return api.ListContext(context.Background())
}

func (api *HooksResource) ListContext(ctx context.Context) (r *HooksListResponse, e error) {
	res, e := api.client.request(ctx, "hooks", string(HttpMethodGet), nil)

	if e != nil {
		return nil, e
	}

	if e := json.Unmarshal([]byte(res), &r); e != nil {
		return nil, e
	}

	return
}
