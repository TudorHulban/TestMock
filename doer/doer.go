package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks github.com/TudorHulban/TestMock/doer Doer

type Doer interface {
	DoSomething(string) error
}
