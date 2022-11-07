package user

import (
	"golang_shane_test/mock"
)

type User struct {
	Person mock.Male
}

func NewUser(p mock.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(a int64) error {
	return u.Person.Get(a)
}
