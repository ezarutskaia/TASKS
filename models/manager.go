package models

import	"errors"

const (
	Open       = "open"
	InProgress = "in progress"
	Done       = "done"
	Closed     = "close"

	Low      = "low"
	Normal   = "normal"
	High     = "high"
	Critical = "critical"
)

var ListStatus = []string{Open, InProgress, Done, Closed}
var ListPriority = []string{Low, Normal, High, Critical}

type Manager struct {
	User User
	AccesLevel int
}

type Task struct {
	Id       int
	Name     string
	UserId   int
	Status   string
	Priority string
}

func isStringExists(slice []string, w string) bool {
	for _, item := range slice {
		if item == w {
			return true
		}
	}
	return false
}

func (m *Manager) CreateTask(name string) (*Task, error) {
	if m.User.hasRole(3) == true {
		return &Task{Id: 1, Name: name, UserId: m.User.Id}, nil
	}
	return &Task{}, errors.New("You don't have roles")

}

func (t *Task) setStatus(s string) (*Task, error) {
	if isStringExists(ListStatus, s) == true {
		t.Status = s
		return t, nil
	}
	return t, errors.New("wrong status")
}

func (t *Task) setPriority(s string) (*Task, error) {
	if isStringExists(ListPriority, s) == true {
		t.Priority = s
		return t, nil
	}
	return t, errors.New("wrong ptiority")
}