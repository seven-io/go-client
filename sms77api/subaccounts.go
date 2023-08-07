package sms77api

import (
	"context"
	"encoding/json"
)

type CreateSubaccountParams struct {
	Action SubaccountsAction `json:"action"`
	Email  string            `json:"email"`
	Name   string            `json:"name"`
}

type CreateSubaccountResponse struct {
	Error      *string     `json:"error"`
	Subaccount *Subaccount `json:"subaccount"`
	Success    bool        `json:"success"`
}

type DeleteSubaccountResponse struct {
	Error   *string `json:"error"`
	Success bool    `json:"success"`
}

type TransferCreditsToSubaccountParams struct {
	Action SubaccountsAction `json:"action"`
	Amount float32           `json:"amount"`
	Id     uint              `json:"id"`
}

type AutoChargeSubaccountParams struct {
	Action    SubaccountsAction `json:"action"`
	Amount    float32           `json:"amount"`
	Id        uint              `json:"id"`
	Threshold float32           `json:"threshold"`
}

type TransferCreditsToSubaccountResponse struct {
	Error   *string `json:"error"`
	Success bool    `json:"success"`
}

type AutoChargeSubaccountResponse struct {
	Error   *string `json:"error"`
	Success bool    `json:"success"`
}

type SubaccountContact struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type SubaccountAutoTopUp struct {
	Amount    *float32 `json:"amount"`
	Threshold *float32 `json:"threshold"`
}

type Subaccount struct {
	AutoTopUp  SubaccountAutoTopUp `json:"auto_topup"`
	Balance    float32             `json:"balance"`
	Company    string              `json:"company"`
	Contact    SubaccountContact   `json:"contact"`
	Id         uint                `json:"id"`
	Username   *string             `json:"username"`
	TotalUsage float32             `json:"total_usage"`
}

type SubaccountsResource resource

type SubaccountsAction string

const (
	SubaccountsActionAutoCharge      SubaccountsAction = "update"
	SubaccountsActionCreate          SubaccountsAction = "create"
	SubaccountsActionDelete          SubaccountsAction = "delete"
	SubaccountsActionRead            SubaccountsAction = "read"
	SubaccountsActionTransferCredits SubaccountsAction = "transfer_credits"
)

func (api *SubaccountsResource) TransferCredits(p TransferCreditsToSubaccountParams) (o TransferCreditsToSubaccountResponse, e error) {
	return api.TransferCreditsContext(context.Background(), p)
}

func (api *SubaccountsResource) TransferCreditsContext(ctx context.Context, p TransferCreditsToSubaccountParams) (o TransferCreditsToSubaccountResponse, e error) {
	p.Action = SubaccountsActionTransferCredits

	r, e := api.client.request(ctx, "subaccounts", "POST", p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}

func (api *SubaccountsResource) AutoCharge(p AutoChargeSubaccountParams) (o AutoChargeSubaccountResponse, e error) {
	return api.AutoChargeContext(context.Background(), p)
}

func (api *SubaccountsResource) AutoChargeContext(ctx context.Context, p AutoChargeSubaccountParams) (o AutoChargeSubaccountResponse, e error) {
	p.Action = SubaccountsActionAutoCharge

	r, e := api.client.request(ctx, "subaccounts", "POST", p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}

func (api *SubaccountsResource) Delete(id uint) (p DeleteSubaccountResponse, e error) {
	return api.DeleteContext(context.Background(), id)
}

func (api *SubaccountsResource) DeleteContext(ctx context.Context, id uint) (o DeleteSubaccountResponse, e error) {
	params := map[string]interface{}{
		"action": SubaccountsActionDelete,
		"id":     id,
	}

	r, e := api.client.request(ctx, "subaccounts", "POST", params)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}

func (api *SubaccountsResource) Create(p CreateSubaccountParams) (o CreateSubaccountResponse, e error) {
	return api.CreateContext(context.Background(), p)
}

func (api *SubaccountsResource) CreateContext(ctx context.Context, p CreateSubaccountParams) (o CreateSubaccountResponse, e error) {
	p.Action = SubaccountsActionCreate

	r, e := api.client.request(ctx, "subaccounts", "POST", p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}

func (api *SubaccountsResource) Read() (o []Subaccount, e error) {
	return api.ReadContext(context.Background())
}

func (api *SubaccountsResource) ReadContext(ctx context.Context) (o []Subaccount, e error) {
	params := map[string]interface{}{
		"action": SubaccountsActionRead,
	}
	r, e := api.client.request(ctx, "subaccounts", "GET", params)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}
