package sms77api

import (
	"fmt"
	"testing"
	"time"
)

func TestSms77API_Hooks(t *testing.T) {
	request := func(params HooksParams) interface{} {
		res, err := client.Hooks.Request(params)

		if err != nil {
			t.Errorf("Hooks() should not return an error, but %s", err)
		}

		if res == nil {
			t.Errorf("Hooks() should return json, but received nil")
		}

		return res
	}

	hooks := request(HooksParams{Action: HooksActionRead}).(*HooksReadResponse)

	if hooks.Success && hooks.Hooks != nil {
		for _, hook := range hooks.Hooks {
			AssertIsLengthy("Created", hook.Created, t)
			AssertIsLengthy("Id", hook.Id, t)
			AssertIsLengthy("TargetUrl", hook.TargetUrl, t)
			AssertInArray("EventType", hook.EventType,
				[...]HookEventType{HookEventTypeSmsStatus, HookEventTypeVoiceStatus, HookEventTypeInboundSms}, t)
			AssertInArray("RequestMethod", hook.RequestMethod,
				[...]HookRequestMethod{HookRequestMethodGet, HookRequestMethodPost}, t)
		}
	}

	subscribed := request(HooksParams{
		Action:        HooksActionSubscribe,
		EventType:     HookEventTypeInboundSms,
		RequestMethod: HookRequestMethodGet,
		TargetUrl:     fmt.Sprintf("https://test.tld/go-client/%d", time.Now().Unix()),
	}).(*HooksSubscribeResponse)

	AssertIsPositive("Id", subscribed.Id, t)

	if true == subscribed.Success {
		subscribed := request(HooksParams{
			Action: HooksActionUnsubscribe,
			Id:     subscribed.Id,
		}).(*HooksUnsubscribeResponse)

		AssertIsTrue("Success", subscribed.Success, t)
	}
}