package models

import "regexp"

type User struct {
	Email string
	Id    int
	Roles []Role
}

type Role struct {
	Name  string
	Value int
}

func (u User) validMail() bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegexPattern, u.Email)
	return matched
}

func (u *User) hasRole(n int) bool {
	for _, item := range u.Roles {
		if item.Value == n {
			return true
		}
	}
	return false
}

func (u *User) isAdmin() bool {
	return u.hasRole(7)
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