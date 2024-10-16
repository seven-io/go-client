package seven

import (
	"context"
	"encoding/json"
)

type ListSubaccountsParams struct {
	ID *uint `json:"id,omitempty"`
}

type CreateSubaccountParams struct {
	Email string `json:"email"`
	Name  string `json:"name"`
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
	Amount float32 `json:"amount"`
	Id     uint    `json:"id"`
}

type AutoChargeSubaccountParams struct {
	Amount    float32 `json:"amount"`
	Id        uint    `json:"id"`
	Threshold float32 `json:"threshold"`
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

func (api *SubaccountsResource) TransferCredits(p TransferCreditsToSubaccountParams) (o TransferCreditsToSubaccountResponse, e error) {
	return api.TransferCreditsContext(context.Background(), p)
}

func (api *SubaccountsResource) TransferCreditsContext(ctx context.Context, p TransferCreditsToSubaccountParams) (o TransferCreditsToSubaccountResponse, e error) {
	r, e := api.client.request(ctx, "subaccounts/transfer_credits", "POST", p)

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
	r, e := api.client.request(ctx, "subaccounts/update", "POST", p)

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
		"id": id,
	}

	r, e := api.client.request(ctx, "subaccounts/delete", "POST", params)

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
	r, e := api.client.request(ctx, "subaccounts/create", "POST", p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}

func (api *SubaccountsResource) Read(p ListSubaccountsParams) (o []Subaccount, e error) {
	return api.ReadContext(context.Background(), p)
}

func (api *SubaccountsResource) ReadContext(ctx context.Context, p ListSubaccountsParams) (o []Subaccount, e error) {
	r, e := api.client.request(ctx, "subaccounts/read", "GET", p)

	if nil != e {
		return
	}

	e = json.Unmarshal([]byte(r), &o)

	return
}
