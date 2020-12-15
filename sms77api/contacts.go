package sms77api

type ContactsResource resource

type ContactsAction string

const (
	ContactsActionDel   ContactsAction = "del"
	ContactsActionRead  ContactsAction = "read"
	ContactsActionWrite ContactsAction = "write"
)

type Contact struct {
	Id    int64  `json:"id"`
	Nick  string `json:"nick"`
	Phone string `json:"empfaenger"`
	EMail string `json:"email"`
}

type ContactsBaseParams struct {
	Action ContactsAction `json:"action"`
}

type ContactsParams struct {
	*ContactsBaseParams
	Json  *bool   `json:"json,omitempty"`
	Id    *int64  `json:"id,omitempty"`
	Nick  *string `json:"nick,omitempty"`
	Phone *string `json:"empfaenger,omitempty"`
	EMail *string `json:"email,omitempty"`
}

type ContactsReadParams struct {
	Json *bool  `json:"json,omitempty"`
	Id   *int64 `json:"id,omitempty"`
}

func (api *ContactsResource) ReadCsv(p ContactsReadParams) (*string, error) {
	type ContactsReadApiParams struct {
		ContactsBaseParams
		ContactsReadParams
	}

	res, err := api.client.request("contacts", "GET", ContactsReadApiParams{
		ContactsBaseParams: ContactsBaseParams{Action: ContactsActionRead},
		ContactsReadParams: p,
	})

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *ContactsResource) Request(p ContactsParams) (*string, error) {
	method := "POST"
	if p.Action == "read" {
		method = "GET"
	}

	res, err := api.client.request("contacts", method, p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
