package seven

import (
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestGroups(t *testing.T) {
	created, e := client.Groups.Create(GroupCreateParams{
		Name: "MyGroup",
	})
	if e != nil {
		t.Errorf(e.Error())
	}

	list, e := client.Groups.List(GroupsListParams{
		Limit:  nil,
		Offset: nil,
	})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.NotEmpty(t, list.Data)

	updateParams := GroupUpdateParams{
		Name: "MyNewGroupName",
	}
	e = client.Groups.Update(created.ID, updateParams)
	if e != nil {
		t.Errorf(e.Error())
	}

	group, e := client.Groups.Get(GroupsGetParams{ID: created.ID})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.NotNil(t, group)
	a.NotEqual(t, created.Name, group.Name)

	deleted, e := client.Groups.Delete(GroupsDeleteParams{ID: group.ID, DeleteContacts: false})
	if e != nil {
		t.Errorf(e.Error())
	}
	a.True(t, deleted.Success)
}
