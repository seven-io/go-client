package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestContacts(t *testing.T) {
	created, e := client.Contacts.Create(ContactCreateParams{
		Avatar:       "",
		Groups:       nil,
		Address:      nil,
		Birthday:     nil,
		City:         nil,
		Email:        nil,
		Firstname:    nil,
		HomeNumber:   nil,
		Lastname:     nil,
		MobileNumber: nil,
		Notes:        nil,
		PostalCode:   nil,
	})
	if e != nil {
		t.Errorf(e.Error())
	}

	list, e := client.Contacts.List(ContactsListParams{
		GroupId:        nil,
		Limit:          nil,
		Offset:         nil,
		OrderBy:        nil,
		OrderDirection: nil,
		Search:         nil,
	})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.NotEmpty(t, list.Data)

	address := "Willestr. 4-6"
	updateParams := ContactUpdateParams{
		Avatar: "",
		Groups: created.Groups,
		//ID:           created.ID,
		Address:      &address,
		Birthday:     nil,
		City:         nil,
		Email:        nil,
		Firstname:    nil,
		HomeNumber:   nil,
		Lastname:     nil,
		MobileNumber: nil,
		Notes:        nil,
		PostalCode:   nil,
	}
	updated, e := client.Contacts.Update(created.ID, updateParams)
	if e != nil {
		t.Errorf(e.Error())
	}
	a.NotEqual(t, created.Properties, updated.Properties)

	contact, e := client.Contacts.Get(ContactsGetParams{ID: updated.ID})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.NotNil(t, contact)

	deleted, e := client.Contacts.Delete(ContactsDeleteParams{ID: contact.ID})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.True(t, deleted.Success)
}
