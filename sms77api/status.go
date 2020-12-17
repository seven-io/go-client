package sms77api

type StatusResource resource

type StatusParams struct {
	MessageId uint64 `json:"msg_id"`
}

const StatusApiCodeInvalidMessageId = "901"

func (api *StatusResource) Post(p StatusParams) (*string, error) {
	res, err := api.client.request("status", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
