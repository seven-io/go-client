package sms77api

import (
	"context"
	"encoding/json"
)

type HookEventType string

const (
	HookEventTypeSmsStatus   HookEventType = "dlr"
	HookEventTypeVoiceStatus HookEventType = "voice_status"
	HookEventTypeInboundSms  HookEventType = "sms_mo"
	HookEventTypeTracking    HookEventType = "tracking"
	HookEventTypeAll         HookEventType = "all"
	HookEventTypeVoiceCall   HookEventType = "voice_call"
)

type HookRequestMethod string

const (
	HookRequestMethodGet  HookRequestMethod = "GET"
	HookRequestMethodJson HookRequestMethod = "JSON"
	HookRequestMethodPost HookRequestMethod = "POST"
)

type HooksAction string

const (
	HooksActionRead        HooksAction = "read"
	HooksActionSubscribe   HooksAction = "subscribe"
	HooksActionUnsubscribe HooksAction = "unsubscribe"
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

type HooksParams struct {
	Action        HooksAction       `json:"action"`
	EventFilter   *string           `json:"event_filter,omitempty"`
	EventType     HookEventType     `json:"event_type,omitempty"`
	Id            int               `json:"id,omitempty"`
	RequestMethod HookRequestMethod `json:"request_method,omitempty"`
	TargetUrl     string            `json:"target_url,omitempty"`
}

type HooksReadResponse struct {
	Success bool   `json:"success"`
	Hooks   []Hook `json:"hooks"`
}

type HooksUnsubscribeResponse struct {
	Success bool `json:"success"`
}

type HooksSubscribeResponse struct {
	Id      int  `json:"id"`
	Success bool `json:"success"`
}

type HooksResource resource

func (api *HooksResource) Request(p HooksParams) (interface{}, error) {
	return api.RequestContext(context.Background(), p)
}

func (api *HooksResource) RequestContext(ctx context.Context, p HooksParams) (interface{}, error) {
	method := "POST"
	if p.Action == HooksActionRead {
		method = "GET"
	}

	res, err := api.client.request(ctx, "hooks", method, p)

	if err != nil {
		return nil, err
	}

	var js interface{}

	switch p.Action {
	case HooksActionRead:
		js = &HooksReadResponse{}
	case HooksActionSubscribe:
		js = &HooksSubscribeResponse{}
	case HooksActionUnsubscribe:
		js = &HooksUnsubscribeResponse{}
	}

	if err := json.Unmarshal([]byte(res), js); err != nil {
		return nil, err
	}

	return js, nil
}
