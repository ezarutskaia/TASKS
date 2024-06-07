package models

import	(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsTaskExists(t *testing.T) {
	tasks := []Task{
		Task{Id: 1, Name: "A", UserId: 3, Status: Open, Priority: Critical},
		Task{Id: 2, Name: "B", UserId: 2, Status: InProgress, Priority: Normal},
		Task{Id: 3, Name: "C", UserId: 1, Status: Closed, Priority: Low},
	}
	onetask := Task{Id: 3, Name: "C", UserId: 1, Status: Closed, Priority: Low}
	expected := true
	actual := IsTaskExists(tasks, onetask)

	assert.Equal(t, expected, actual)
}

func TestMatchTask(t *testing.T) {
	filter := TaskFilter{Type: "status", Value: Open}
	task := Task{Id: 5, Name: "F", UserId: 2, Status: Closed, Priority: Critical}
	expected := false
	actual := filter.MatchTask(task)

	assert.Equal(t, expected, actual)
}

func TestFilterTaskSlice(t *testing.T) {
	tasks := []Task{
		Task{Id: 3, Name: "C", UserId: 2, Status: InProgress, Priority: Low},
		Task{Id: 4, Name: "D", UserId: 2, Status: InProgress, Priority: High},
		Task{Id: 5, Name: "F", UserId: 2, Status: Open, Priority: Critical},
	}
	filters := []TaskFilter{
		TaskFilter{Type: FilterStatus, Value: Open},
		TaskFilter{Type: FilterPriority, Value: Critical},
	}
	expected := []Task {Task{Id: 5, Name: "F", UserId: 2, Status: Open, Priority: Critical},}
	actual := FilterTaskSlice(tasks, filters)

	assert.Equal(t, expected, actual)
}