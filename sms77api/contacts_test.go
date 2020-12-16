package sms77api

import (
	"fmt"
	"strings"
	"testing"
)

var createdId uint64

func assertContact(c Contact, t *testing.T) {
	if 0 == toUint64(c.Id) {
		t.Error("Every Contact must have a positive ID")
	}
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

	AssertEquals("code", actual, string(*getExpectedContactCode(err)), t)
}

func getDeleteParams() ContactsDeleteParams {
	if 0 == createdId {
		o, _ := client.Contacts.CreateJson()

		createdId = o.Id
	}

	return ContactsDeleteParams{createdId}
}

func isValidCode(needle ContactsWriteCode) bool {
	return InArray(needle, [...]ContactsWriteCode{ContactsWriteCodeUnchanged, ContactsWriteCodeChanged})
}

func prepareEdit() (Contact, string) {
	contacts, _ := client.Contacts.ReadJson(ContactsReadParams{})
	contact := contacts[0]

	return contact, fmt.Sprintf("%sXXX", contact.Nick)
}

func TestContacts_CreateCsv(t *testing.T) {
	text, err := client.Contacts.CreateCsv()

	if nil == err {
		lines := strings.Split(text, "\n")

		if AssertEquals("api_code", lines[0], "152", t) {
			id := toUint64(lines[1])
			if AssertIsPositive("created_id", id, t) {
				createdId = id
			}
		} else {
			AssertEquals("linesCount", len(lines), 1, t)
		}

		if 0 != createdId {
			TestContacts_DeleteCsv(t)
		}
	} else {
		AssertEquals("response", text, nil, t)
	}
}

func TestContacts_CreateJson(t *testing.T) {
	o, e := client.Contacts.CreateJson()

	if nil == e {
		AssertIsTrue("valid_return_code", isValidCode(o.Return), t)

		if ContactsWriteCodeChanged == o.Return {
			createdId = o.Id

			TestContacts_DeleteJson(t)
		} else {
			AssertEquals("contact_id", 0, nil, t)
		}
	} else {
		AssertEquals("response", o, nil, t)
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
		AssertIsTrue("valid_code", isValidCode(writeCode), t)

		if ContactsWriteCodeChanged != writeCode {
			expectedNick = contact.Nick
		}

		contacts, _ := client.Contacts.ReadJson(ContactsReadParams{toUint64(contact.Id)})
		AssertEquals("contact.Nick", contacts[0].Nick, expectedNick, t)
	} else {
		AssertIsNil("response", code, t)
	}
}

func TestContacts_EditJson(t *testing.T) {
	contact, expectedNick := prepareEdit()

	obj, err := client.Contacts.EditJson(ContactEditParams{Id: contact.Id, Nick: expectedNick})

	if nil == err {
		if ContactsWriteCodeChanged != obj.Return {
			expectedNick = contact.Nick
		}

		contacts, _ := client.Contacts.ReadJson(ContactsReadParams{toUint64(contact.Id)})
		AssertEquals("contact.Nick", contacts[0].Nick, expectedNick, t)
	} else {
		AssertIsNil("response", obj, t)
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
		AssertEquals("csv", csv, nil, t)
	}
}

func TestContacts_ReadJson(t *testing.T) {
	array, err := client.Contacts.ReadJson(ContactsReadParams{})

	if nil == err {
		for _, contact := range array {
			assertContact(contact, t)
		}
	} else {
		AssertEquals("contacts", array, nil, t)
	}
}
