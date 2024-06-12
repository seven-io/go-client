package seven

import (
	"context"
	"encoding/json"
	"fmt"
)

type ContactsResource resource

type ContactsList struct {
	Data           []Contact      `json:"data"`
	PagingMetadata PagingMetadata `json:"pagingMetadata"`
}

type PagingMetadata struct {
	Count   uint `json:"count"`
	HasMore bool `json:"has_more"`
	Limit   uint `json:"limit"`
	Offset  uint `json:"offset"`
	Total   uint `json:"total"`
}

type Contact struct {
	Avatar     string            `json:"avatar"`
	Created    string            `json:"created"`
	Groups     []uint            `json:"groups"`
	ID         uint              `json:"id"`
	Initials   ContactInitials   `json:"initials"`
	Properties ContactProperties `json:"properties"`
	Validation ContactValidation `json:"validation"`
}

type ContactValidation struct {
	State     *string `json:"state"`
	Timestamp *string `json:"timestamp"`
}

type ContactInitials struct {
	Color    string `json:"color"`
	Initials string `json:"initials"`
}

type ContactProperties struct {
	Address      *string `json:"address"`
	Birthday     *string `json:"birthday"`
	City         *string `json:"city"`
	Email        *string `json:"email"`
	Firstname    *string `json:"firstname"`
	Fullname     *string `json:"fullname"`
	HomeNumber   *string `json:"home_number"`
	Lastname     *string `json:"lastname"`
	MobileNumber *string `json:"mobile_number"`
	Notes        *string `json:"notes"`
	PostalCode   *string `json:"postal_code"`
}

type ContactCreateParams struct {
	Avatar string `json:"avatar,omitempty"`
	Groups []uint `json:"groups,omitempty"`

	Address      *string `json:"address,omitempty"`
	Birthday     *string `json:"birthday,omitempty"`
	City         *string `json:"city,omitempty"`
	Email        *string `json:"email,omitempty"`
	Firstname    *string `json:"firstname,omitempty"`
	HomeNumber   *string `json:"home_number,omitempty"`
	Lastname     *string `json:"lastname,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	Notes        *string `json:"notes,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
}

type ContactUpdateParams struct {
	Avatar string `json:"avatar,omitempty"`
	Groups []uint `json:"groups,omitempty"`
	//ID     uint   `json:"id"`

	Address      *string `json:"address,omitempty"`
	Birthday     *string `json:"birthday,omitempty"`
	City         *string `json:"city,omitempty"`
	Email        *string `json:"email,omitempty"`
	Firstname    *string `json:"firstname,omitempty"`
	HomeNumber   *string `json:"home_number,omitempty"`
	Lastname     *string `json:"lastname,omitempty"`
	MobileNumber *string `json:"mobile_number,omitempty"`
	Notes        *string `json:"notes,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
}

type ContactsDeleteParams = struct {
	ID uint `json:"id"`
}

type ContactsDeleteResponse = struct {
	Success bool `json:"success"`
}

type ContactsGetParams = struct {
	ID uint `json:"id"`
}

type ContactsListParams = struct {
	GroupId        *uint   `json:"group_id,omitempty"`
	Limit          *uint   `json:"limit,omitempty"`
	Offset         *uint   `json:"offset,omitempty"`
	OrderBy        *string `json:"order_by,omitempty"`
	OrderDirection *string `json:"order_direction,omitempty"`
	Search         *string `json:"search,omitempty"`
}

func (api *ContactsResource) Get(p ContactsGetParams) (c Contact, e error) {
	return api.GetContext(context.Background(), p)
}

func (api *ContactsResource) GetContext(ctx context.Context, p ContactsGetParams) (c Contact, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("contacts/%d", p.ID), string(HttpMethodGet), nil)

	if nil != e {
		return
	}

	json.Unmarshal([]byte(s), &c)

	return
}

func (api *ContactsResource) List(p ContactsListParams) (a ContactsList, e error) {
	return api.ListContext(context.Background(), p)
}

func (api *ContactsResource) ListContext(ctx context.Context, p ContactsListParams) (a ContactsList, e error) {
	s, e := api.client.request(ctx, "contacts", string(HttpMethodGet), p)

	if nil != e {
		return
	}

	json.Unmarshal([]byte(s), &a)

	return
}

func (api *ContactsResource) Create(p ContactCreateParams) (c Contact, e error) {
	return api.CreateContext(context.Background(), p)
}

func (api *ContactsResource) CreateContext(ctx context.Context, p ContactCreateParams) (c Contact, e error) {
	s, e := api.client.request(ctx, "contacts", string(HttpMethodPost), p)

	e = json.Unmarshal([]byte(s), &c)

	return
}

func (api *ContactsResource) Delete(p ContactsDeleteParams) (o ContactsDeleteResponse, e error) {
	return api.DeleteContext(context.Background(), p)
}

func (api *ContactsResource) DeleteContext(ctx context.Context, p ContactsDeleteParams) (o ContactsDeleteResponse, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("contacts/%d", p.ID), string(HttpMethodDelete), p)

	e = json.Unmarshal([]byte(s), &o)

	return
}

func (api *ContactsResource) Update(id uint, p ContactUpdateParams) (c Contact, e error) {
	return api.UpdateContext(context.Background(), id, p)
}

func (api *ContactsResource) UpdateContext(ctx context.Context, id uint, p ContactUpdateParams) (c Contact, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("contacts/%d", id), string(HttpMethodPatch), p)

	e = json.Unmarshal([]byte(s), &c)

	return
}
