package sms77api

type StatusParams struct {
	MessageId int64 `json:"msg_id"`
}

type StatusResource resource

func (api *StatusResource) Post(p StatusParams) (*string, error) {
	res, err := api.client.request("status", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
