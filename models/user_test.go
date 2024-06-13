package models

import	(
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestValidMail(t *testing.T) {
	testUser1 := &User{Email: "xyz@nextmail.com"}
	testUser2 := &User{Email: "xyz%nextmail.com"}
	expected1 := true
	actual1 := testUser1.ValidMail()
	expected2 := false
	actual2 := testUser2.ValidMail()

	assert.Equal(t, strconv.FormatBool(expected1), strconv.FormatBool(actual1))
	assert.Equal(t, strconv.FormatBool(expected2), strconv.FormatBool(actual2))
}

func TestHasRole(t *testing.T) {
	testUser := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3},}}
	n := 3
	expected := true
	actual := testUser.HasRole(n)

	assert.Equal(t, expected, actual)
}

func TestIsAdmin(t *testing.T) {
	testUser := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3},}}
	expected := false
	actual := testUser.IsAdmin()

	assert.Equal(t, expected, actual)
}

func TestAddRole(t *testing.T) {
	testUser := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3},}}
	testRole := Role{Id: 2, Name:"Admin", Value:7}
	expected := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3}, Role{Id: 2, Name:"Admin", Value:7},}}
	testUser.AddRole(testRole)

	assert.Equal(t, expected, testUser)
}

func TestRevokeRole(t *testing.T) {
	testUser := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3}, Role{Id: 2, Name:"Admin", Value:7},}}
	testRole := Role{Id: 2, Name:"Admin", Value:7}
	expected := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3},}}
	testUser.RevokeRole(testRole)

	assert.Equal(t, expected, testUser)
}