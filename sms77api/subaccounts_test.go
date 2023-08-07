package sms77api

import (
	"fmt"
	a "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSubaccountsResource_TransferCredits(t *testing.T) {
	res, err := client.Subaccounts.TransferCredits(TransferCreditsToSubaccountParams{
		Amount: -0.0,
		Id:     0,
	})

	if err != nil {
		a.Nil(t, res)
	} else {
		a.False(t, res.Success)
		a.NotEmpty(t, res.Error)
		a.NotNil(t, res.Error)
	}

	// success

	res2, _ := client.Subaccounts.Create(CreateSubaccountParams{
		Email: fmt.Sprintf("tommy_tester_%v@seven.dev", time.Now().Unix()),
		Name:  "Tommy Tester",
	})
	res3, err3 := client.Subaccounts.TransferCredits(TransferCreditsToSubaccountParams{
		Amount: 1.1,
		Id:     res2.Subaccount.Id,
	})

	if err3 != nil {
		a.Nil(t, res3)
		client.Subaccounts.Delete(res2.Subaccount.Id)
	} else {
		if res3.Success {
			a.Nil(t, res3.Error)
		} else {
			a.NotNil(t, res3.Error)
		}
	}
}

func TestSubaccountsResource_AutoCharge(t *testing.T) {
	res, err := client.Subaccounts.AutoCharge(AutoChargeSubaccountParams{
		Amount:    -0.0,
		Id:        0,
		Threshold: -0.0,
	})

	if err != nil {
		a.Nil(t, res)
	} else {
		a.False(t, res.Success)
		a.NotEmpty(t, res.Error)
		a.NotNil(t, res.Error)
	}

	// success

	res2, _ := client.Subaccounts.Create(CreateSubaccountParams{
		Email: fmt.Sprintf("tommy_tester_%v@seven.dev", time.Now().Unix()),
		Name:  "Tommy Tester",
	})
	res3, err3 := client.Subaccounts.AutoCharge(AutoChargeSubaccountParams{
		Amount:    1.1,
		Id:        res2.Subaccount.Id,
		Threshold: 2.2,
	})

	if err3 != nil {
		a.Nil(t, res3)
		client.Subaccounts.Delete(res2.Subaccount.Id)
	} else {
		if res3.Success {
			a.Nil(t, res3.Error)
		} else {
			a.NotNil(t, res3.Error)
		}
	}
}

func TestSubaccountsResource_Create(t *testing.T) {
	res, err := client.Subaccounts.Create(CreateSubaccountParams{
		Email: "",
		Name:  "",
	})

	if err != nil {
		a.Nil(t, res)
	} else {
		a.Nil(t, res.Subaccount)
		a.NotEmpty(t, res.Error)
		a.False(t, res.Success)
	}

	// success

	p := CreateSubaccountParams{
		Email: fmt.Sprintf("tommy_tester_%v@seven.dev", time.Now().Unix()),
		Name:  "Tommy Tester",
	}
	res2, err2 := client.Subaccounts.Create(p)

	if err2 != nil {
		a.Nil(t, res2)
	} else {
		a.Empty(t, res2.Error)
		a.NotNil(t, res2.Subaccount)
		a.True(t, res2.Success)

		a.Equal(t, float32(0.0), *res2.Subaccount.AutoTopUp.Amount)
		a.Equal(t, float32(0.0), *res2.Subaccount.AutoTopUp.Threshold)

		a.Equal(t, p.Name, res2.Subaccount.Contact.Name)
		a.Equal(t, p.Email, res2.Subaccount.Contact.Email)

		//goland:noinspection GoUnhandledErrorResult
		client.Subaccounts.Delete(res2.Subaccount.Id)
	}
}

func TestSubaccountsResource_Read(t *testing.T) {
	res, err := client.Subaccounts.Read()

	if err != nil {
		a.Nil(t, res)
	} else {
		for _, subaccount := range res {
			a.GreaterOrEqual(t, subaccount.Balance, float32(0.0))
			a.Greater(t, subaccount.Id, uint(0))
		}
	}
}
