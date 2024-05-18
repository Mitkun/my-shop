package userdomain

import (
	"github.com/google/uuid"
	"strings"
)

type User struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	address     string
	phoneNumber string
	email       string
	password    string
	salt        string
	status      string
	role        Role
}

func NewUser(id uuid.UUID, firstName, lastName, address, phoneNumber, email, password, salt, status string, role Role) (*User, error) {
	return &User{id: id, firstName: firstName, lastName: lastName, address: address, phoneNumber: phoneNumber, email: email, password: password, salt: salt, status: status, role: role}, nil
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Address() string {
	return u.address
}

func (u *User) PhoneNumber() string {
	return u.phoneNumber
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Salt() string {
	return u.salt
}

func (u *User) Role() Role {
	return u.role
}
func (u User) Status() string {
	return u.status
}

const (
	RoleUser Role = iota
	RoleAdmin
)

type Role int

func (r Role) String() string {
	return [2]string{"user", "admin"}[r]
}

func GetRole(s string) Role {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "admin":
		return RoleAdmin
	default:
		return RoleUser
	}
}
