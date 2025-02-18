package seven

import (
	"testing"

	a "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroups(t *testing.T) {
	created, e := client.Groups.Create(GroupCreateParams{
		Name: "MyGroup",
	})
	require.NoError(t, e)

	list, e := client.Groups.List(GroupsListParams{
		Limit:  nil,
		Offset: nil,
	})
	if a.NoError(t, e) {
		a.NotEmpty(t, list.Data)
		a.Contains(t, list.Data, created)
	}

	updateParams := GroupUpdateParams{
		Name: "MyNewGroupName",
	}
	e = client.Groups.Update(created.ID, updateParams)
	a.NoError(t, e)

	group, e := client.Groups.Get(GroupsGetParams{ID: created.ID})
	if a.NoError(t, e) {
		a.NotNil(t, group)
		a.NotEqual(t, created.Name, group.Name)
		a.Equal(t, updateParams.Name, group.Name)
		a.Equal(t, created.ID, group.ID)
	}

	deleted, e := client.Groups.Delete(GroupsDeleteParams{ID: group.ID, DeleteContacts: false})
	if a.NoError(t, e) {
		a.True(t, deleted.Success)
	}
}
