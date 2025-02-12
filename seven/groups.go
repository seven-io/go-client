package seven

import (
	"context"
	"encoding/json"
	"fmt"
)

type GroupsResource resource

type GroupsList struct {
	Data           []Group        `json:"data"`
	PagingMetadata PagingMetadata `json:"pagingMetadata"`
}

type Group struct {
	Name         string `json:"name"`
	Created      string `json:"created"`
	MembersCount uint   `json:"members_count"`
	ID           uint   `json:"id"`
}

type GroupOrderParams struct {
	Number          string          `json:"number"`
	PaymentInterval PaymentInterval `json:"payment_interval"`
}

type GroupCreateParams struct {
	// Name of the group
	Name string `json:"name"`
}

type GroupUpdateParams struct {
	// Name is the new name of the group.
	Name string `json:"name"`
}

type GroupsDeleteParams = struct {
	ID             uint `json:"id"`
	DeleteContacts bool `json:"delete_contacts,omitempty"`
}

type GroupsDeleteResponse = struct {
	Success bool `json:"success"`
}

type GroupsGetParams = struct {
	ID uint `json:"id"`
}

type GroupsListParams = struct {
	Limit  *uint `json:"limit,omitempty"`
	Offset *uint `json:"offset,omitempty"`
}

func (api *GroupsResource) Get(p GroupsGetParams) (c Group, e error) {
	return api.GetContext(context.Background(), p)
}

func (api *GroupsResource) GetContext(ctx context.Context, p GroupsGetParams) (c Group, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("groups/%d", p.ID), string(HttpMethodGet), nil)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(s), &c)

	return
}

func (api *GroupsResource) List(p GroupsListParams) (a GroupsList, e error) {
	return api.ListContext(context.Background(), p)
}

func (api *GroupsResource) ListContext(ctx context.Context, p GroupsListParams) (a GroupsList, e error) {
	s, e := api.client.request(ctx, "groups", string(HttpMethodGet), p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(s), &a)

	return
}

func (api *GroupsResource) Create(p GroupCreateParams) (c Group, e error) {
	return api.CreateContext(context.Background(), p)
}

func (api *GroupsResource) CreateContext(ctx context.Context, p GroupCreateParams) (c Group, e error) {
	s, e := api.client.request(ctx, "groups", string(HttpMethodPost), p)
	if e != nil {
		return
	}

	e = json.Unmarshal([]byte(s), &c)
	return
}

func (api *GroupsResource) Delete(p GroupsDeleteParams) (o GroupsDeleteResponse, e error) {
	return api.DeleteContext(context.Background(), p)
}

func (api *GroupsResource) DeleteContext(ctx context.Context, p GroupsDeleteParams) (o GroupsDeleteResponse, e error) {
	s, e := api.client.request(ctx, fmt.Sprintf("groups/%d", p.ID), string(HttpMethodDelete), p)
	if e != nil {
		return
	}

	e = json.Unmarshal([]byte(s), &o)

	return
}

func (api *GroupsResource) Update(id uint, p GroupUpdateParams) (e error) {
	return api.UpdateContext(context.Background(), id, p)
}

func (api *GroupsResource) UpdateContext(ctx context.Context, id uint, p GroupUpdateParams) (e error) {
	_, e = api.client.request(ctx, fmt.Sprintf("groups/%d", id), string(HttpMethodPatch), p)

	return
}
