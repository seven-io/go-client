package seven

import (
	"fmt"
	a "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHooks(t *testing.T) {
	subscribed, err := client.Hooks.Subscribe(HooksSubscribeParams{
		EventFilter:   "",
		EventType:     HookEventTypeInboundSms,
		RequestMethod: HookRequestMethodGet,
		TargetUrl:     fmt.Sprintf("https://test.tld/go-client/%d", time.Now().Unix()),
	})
	if err != nil {
		t.Errorf(err.Error())
	}

	a.Greater(t, subscribed.Id, uint(0))
	a.True(t, subscribed.Success)

	hooks, err := client.Hooks.List()
	if err != nil {
		t.Errorf(err.Error())
	}
	a.True(t, hooks.Success)
	a.NotEmpty(t, hooks.Hooks)

	unsubscribed, err := client.Hooks.Unsubscribe(subscribed.Id)
	if err != nil {
		t.Errorf(err.Error())
	}
	a.True(t, unsubscribed.Success)
}
