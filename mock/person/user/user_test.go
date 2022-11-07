package user

import (
	"golang_shane_test/mock/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var a int64 = 1
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(a).Return(nil),
	)

	user := NewUser(mockMale)
	err := user.GetUserInfo(a)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}
