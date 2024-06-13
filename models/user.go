package models

import (
	"regexp"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id    int
	Email string `gorm:"unique"`
	Roles []Role `gorm:"many2many:user_roles;"`
	Tasks []Task
}

type Role struct {
	gorm.Model
	Id int
	Name  string
	Value int
}

func (u User) ValidMail() bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegexPattern, u.Email)
	return matched
}

func (u *User) HasRole(n int) bool {
	for _, item := range u.Roles {
		if item.Value == n {
			return true
		}
	}
	return false
}

func (u *User) IsAdmin() bool {
	return u.HasRole(7)
}

func (u *User) AddRole(r Role) {
	var isRoleExists = false
	for _, item := range u.Roles {
		if item.Name == r.Name {
			isRoleExists = true
		}
	}
	if isRoleExists == false {
		u.Roles = append(u.Roles, r)
	}
}

func (u *User) RevokeRole(r Role) {
	var isRoleExists = false
	var i = 0
	for ind, item := range u.Roles {
		if item.Name == r.Name {
			isRoleExists = true
			i = ind
		}
	}
	if isRoleExists == true {
		u.Roles[i] = u.Roles[len(u.Roles)-1]
		u.Roles = u.Roles[:len(u.Roles)-1]
	}
}