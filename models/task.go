package models

import	(
	"errors"
	"gorm.io/gorm"
)

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

type Task struct {
	gorm.Model
	Id       int `gorm:"primaryKey"`
	Name     string
	UserID   int
	Status   string
	Priority string
}

func IsStringExists(slice []string, w string) bool {
	for _, item := range slice {
		if item == w {
			return true
		}
	}
	return false
}

func (u *User) CreateTask(name string) (*Task, error) {
	if u.HasRole(3) == true {
		return &Task{Id: 1, Name: name, UserID: u.Id}, nil
	}
	return &Task{}, errors.New("You don't have roles")

}

func (t *Task) SetStatus(s string) (*Task, error) {
	if IsStringExists(ListStatus, s) == true {
		t.Status = s
		return t, nil
	}
	return t, errors.New("wrong status")
}

func (t *Task) SetPriority(s string) (*Task, error) {
	if IsStringExists(ListPriority, s) == true {
		t.Priority = s
		return t, nil
	}
	return t, errors.New("wrong ptiority")
}