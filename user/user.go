package user

import "github.com/TudorHulban/TestMock/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) PrintMsg(m string) error {
	return u.Doer.DoSomething(m)
}
