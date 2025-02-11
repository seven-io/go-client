package sms77api

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	a "github.com/stretchr/testify/assert"
)

var createdId uint64

func assertContact(c Contact, t *testing.T) {
	a.Greater(t, toUint(c.Id, 64), uint64(0))
}

func assertDeletion(actual string, err error, t *testing.T) {
	getExpectedContactCode := func(err error) *ContactsWriteCode {
		if nil != err {
			return nil
		}

		var expected ContactsWriteCode

		if 0 == createdId {
			expected = ContactsWriteCodeUnchanged
		} else {
			expected = ContactsWriteCodeChanged

			createdId = 0
		}

		return &expected
	}

	a.Equal(t, string(*getExpectedContactCode(err)), actual)
}

func getDeleteParams() ContactsDeleteParams {
	if 0 == createdId {
		o, _ := client.Contacts.CreateJson()

		createdId = o.Id
	}

	return ContactsDeleteParams{createdId}
}

func isValidCode(needle ContactsWriteCode) bool {
	validCodes := []ContactsWriteCode{ContactsWriteCodeUnchanged, ContactsWriteCodeChanged}
	return slices.Contains(validCodes, needle)
}

func prepareEdit() (Contact, string) {
	contacts, _ := client.Contacts.ReadJson(ContactsReadParams{})

	if 0 == len(contacts) {
		c, _ := client.Contacts.CreateJson()
		contacts, _ = client.Contacts.ReadJson(ContactsReadParams{c.Id})
	}

	contact := contacts[0]

	return contact, fmt.Sprintf("%sXXX", contact.Nick)
}

func TestContacts_CreateCsv(t *testing.T) {
	text, err := client.Contacts.CreateCsv()

	if nil == err {
		lines := strings.Split(text, "\n")

		if a.Equal(t, "152", lines[0]) {
			id := toUint(lines[1], 64)
			if a.Greater(t, id, uint64(0)) {
				createdId = id
			}
		} else {
			a.Equal(t, 1, len(lines))
		}

		if 0 != createdId {
			TestContacts_DeleteCsv(t)
		}
	} else {
		a.Equal(t, nil, text)
	}
}

func TestContacts_CreateJson(t *testing.T) {
	o, e := client.Contacts.CreateJson()

	if nil == e {
		a.True(t, isValidCode(o.Return))

		if ContactsWriteCodeChanged == o.Return {
			createdId = o.Id

			TestContacts_DeleteJson(t)
		} else {
			a.Equal(t, ContactsWriteCodeUnchanged, o.Return)
		}
	} else {
		a.Equal(t, nil, o)
	}
}

func TestContacts_DeleteCsv(t *testing.T) {
	code, err := client.Contacts.DeleteCsv(getDeleteParams())

	assertDeletion(code, err, t)
}

func TestContacts_DeleteJson(t *testing.T) {
	code, err := client.Contacts.DeleteJson(getDeleteParams())

	var actual ContactsWriteCode

	if nil == err {
		actual = code.Return
	}

	assertDeletion(string(actual), err, t)
}

func TestContacts_EditCsv(t *testing.T) {
	contact, expectedNick := prepareEdit()

	code, err := client.Contacts.EditCsv(ContactEditParams{Id: contact.Id, Nick: contact.Nick})

	if nil == err {
		writeCode := ContactsWriteCode(code)
		a.True(t, isValidCode(writeCode))

		if ContactsWriteCodeChanged != writeCode {
			expectedNick = contact.Nick
		}

		contacts, _ := client.Contacts.ReadJson(ContactsReadParams{toUint(contact.Id, 64)})
		a.Equal(t, expectedNick, contacts[0].Nick)
	} else {
		a.Nil(t, code)
	}
}

func TestContacts_EditJson(t *testing.T) {
	contact, expectedNick := prepareEdit()

	obj, err := client.Contacts.EditJson(ContactEditParams{Id: contact.Id, Nick: expectedNick})

	if nil == err {
		if ContactsWriteCodeChanged != obj.Return {
			expectedNick = contact.Nick
		}

		contacts, _ := client.Contacts.ReadJson(ContactsReadParams{toUint(contact.Id, 64)})
		a.Equal(t, expectedNick, contacts[0].Nick)
	} else {
		a.Nil(t, obj)
	}
}

func TestContacts_ReadCsv(t *testing.T) {
	toStruct := func(c string) Contact {
		c = strings.TrimSpace(c)
		arr := strings.Split(c, ";")

		return Contact{
			Id:    strings.ReplaceAll(arr[0], "\"", ""),
			Nick:  strings.ReplaceAll(arr[1], "\"", ""),
			Phone: strings.ReplaceAll(arr[2], "\"", ""),
		}
	}

	csv, err := client.Contacts.ReadCsv(ContactsReadParams{})

	if nil == err {
		for _, csv := range strings.Split(strings.TrimSpace(csv), "\n") {

			assertContact(toStruct(csv), t)
		}
	} else {
		a.Equal(t, nil, csv)
	}
}

func TestContacts_ReadJson(t *testing.T) {
	array, err := client.Contacts.ReadJson(ContactsReadParams{})

	if nil == err {
		for _, contact := range array {
			assertContact(contact, t)
		}
	} else {
		a.Equal(t, nil, array)
	}
}
