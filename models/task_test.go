package models

import	(
	"testing"
	"errors"
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

	assert.Equal(t, strconv.FormatBool(expected1), strconv.FormatBool(actual1))
	assert.Equal(t, strconv.FormatBool(expected2), strconv.FormatBool(actual2))
}

func TestCreateTask(t *testing.T) {
	name := "testName"
	testUser := &User{Email: "xyz@mail.com", Id: 3, Roles: []Role{Role{Id: 1, Name:"Worker", Value: 3},}}
	expected := &Task{Id: 1, Name: "testName", UserId: 3}
	actual, err := testUser.CreateTask(name)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSetStatus(t *testing.T) {
	testTask := &Task{Id: 1, Name: "testName", UserId: 3}
	testStatus := "in progress"
	expected := &Task{Id: 1, Name: "testName", UserId: 3, Status: "in progress"}
	actual, err := testTask.SetStatus(testStatus)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestSetPriority(t *testing.T) {
	testTask := &Task{Id: 1, Name: "testName", UserId: 3}
	testPriority := "ok"
	expected := errors.New("wrong ptiority")
	_, err := testTask.SetPriority(testPriority)

	assert.Equal(t, expected, err)
}