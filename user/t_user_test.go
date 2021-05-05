package user_test

import (
	"testing"

	"github.com/TudorHulban/TestMock/mocks"
	"github.com/TudorHulban/TestMock/user"
	gomock "github.com/golang/mock/gomock"
)

func TestDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{
		Doer: mockDoer,
	}

	// Expect Do to be called once with "xxx" as parameter, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething("xxx").Return(nil).Times(1)

	testUser.PrintMsg()
}
