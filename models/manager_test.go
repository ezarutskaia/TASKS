package models

import	(
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)


func TestIsStringExists(t *testing.T) {
	
	slice := []string{"str1","str2","str3"}
	text1 := "str2"
	text2 := "str4"
	expected1 := true
	actual1 := IsStringExists(slice, text1)
	expected2 := false
	actual2 := IsStringExists(slice, text2)

	// if expected1 != actual1 {
    // 	t.Errorf("Result was incorrect, got: %s, want: %s.", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	// }
	// if expected2 != actual2 {
    // 	t.Errorf("Result was incorrect, got: %s, want: %s.", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	// }
	assert.Equal(t, strconv.FormatBool(expected1), strconv.FormatBool(actual1))
	assert.Equal(t, strconv.FormatBool(expected2), strconv.FormatBool(actual2))

}

func TestCreateTask(t *testing.T) {
	name := "testName"
	//testRole := &Role{"Worker", 3}
	testUser := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{"Worker", 3},}}
	testManager := &Manager{User: *testUser, AccesLevel: 3}
	expected := &Task{Id: 1, Name: "testName", UserId: 3}
	actual, err := testManager.CreateTask(name)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

