package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func testStatusAssert(s *Status, t *testing.T) {
	a.NotEmpty(t, s.Code)
	a.NotEmpty(t, s.DateTime)
}

func testStatusText(messageId uint64, t *testing.T) *string {
	status, err := client.Status.Text(StatusParams{MessageId: messageId})

	if nil != err {
		a.Nil(t, status)
	}

	return status
}

func testStatusJson(messageId uint64, t *testing.T) *Status {
	status, err := client.Status.Json(StatusParams{MessageId: messageId})

	if nil != err {
		a.Nil(t, status)
	}

	return status
}

func testStatusGetId() uint64 {
	journals, _ := client.Journal.Outbound(&JournalParams{})
	var id string

	if 0 == len(journals) {
		sms, _ := client.Sms.Json(SmsBaseParams{To: VinTelekom, Text: "HI"})
		id = *sms.Messages[0].Id
	} else {
		id = journals[0].Id
	}

	return toUint(id, 64)
}

func TestStatusResource_Text(t *testing.T) {
	s, _ := makeStatus(testStatusText(testStatusGetId(), t))

	testStatusAssert(s, t)
}

func TestStatusResource_Text_Fail(t *testing.T) {
	a.Equal(t, StatusApiCodeInvalidMessageId, *testStatusText(0, t))
}

func TestStatusResource_Json(t *testing.T) {
	testStatusAssert(testStatusJson(testStatusGetId(), t), t)
}

func TestStatusResource_Json_Fail(t *testing.T) {
	a.Nil(t, testStatusJson(0, t))
}
