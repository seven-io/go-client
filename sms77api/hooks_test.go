package sms77api

import (
	"fmt"
	a "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func hooks(p HooksParams, t *testing.T) interface{} {
	res, err := client.Hooks.Request(p)

	if err != nil {
		a.Nil(t, res)
	}

	return res
}

func TestHooksRead(t *testing.T) {
	hooks := hooks(HooksParams{Action: HooksActionRead}, t).(*HooksReadResponse)

	if hooks.Success && hooks.Hooks != nil {
		for _, hook := range hooks.Hooks {
			a.Greater(t, len(hook.Created), 0)
			a.Greater(t, toUint(hook.Id, 64), uint64(0))
			a.Greater(t, len(hook.TargetUrl), 0)
			a.Contains(t, [...]HookEventType{
				HookEventTypeSmsStatus,
				HookEventTypeVoiceStatus,
				HookEventTypeInboundSms,
				HookEventTypeTracking,
				HookEventTypeAll,
				HookEventTypeVoiceCall,
			},
				hook.EventType)
			a.Contains(t, [...]HookRequestMethod{HookRequestMethodGet, HookRequestMethodJson, HookRequestMethodPost},
				hook.RequestMethod)
		}
	}
}

func TestHooksSubscribeAndUnsubscribe(t *testing.T) {
	subscribed := hooks(HooksParams{
		Action:        HooksActionSubscribe,
		EventType:     HookEventTypeInboundSms,
		RequestMethod: HookRequestMethodGet,
		TargetUrl:     fmt.Sprintf("https://test.tld/go-client/%d", time.Now().Unix()),
	}, t).(*HooksSubscribeResponse)

	a.Greater(t, subscribed.Id, 0)

	if true == subscribed.Success {
		subscribed := hooks(HooksParams{
			Action: HooksActionUnsubscribe,
			Id:     subscribed.Id,
		}, t).(*HooksUnsubscribeResponse)

		a.True(t, subscribed.Success)
	}
}
