package car

import (
	gotest "golang_shane_test/car/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

//	func Test_DOT(t *testing.T) {
//		if i, e := DOT(6, 2); i != 12 || e != nil { //try a unit test on function
//			t.Error("乘法函式測試沒通過") // 如果不是如預期的那麼就報錯
//		}
//	}
func Test_UseDivision(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	// mockMale := mock.NewMockMale(ctl)
	mockMale := gotest.NewMockdataInterface(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Division(float64(3), float64(2)).Return(1.5, nil),
	)
	// user := Division(mockMale)
	c := UseData{mockMale}
	f, err := c.UseDivision(1, 2)
	if err != nil {
		t.Error("err, want error:", err)
	}

	if f != 4.5 {
		t.Error("除法錯誤")
	}

}
func Test_UseDot(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	// mockMale := mock.NewMockMale(ctl)
	mockMale := gotest.NewMockdataInterface(ctl)
	gomock.InOrder(
		mockMale.EXPECT().DOT(float64(3), float64(2)).Return(float64(6), nil),
	)
	c := UseData{mockMale}
	f, err := c.UseDot(3, 2)
	if err != nil {
		t.Error("err, want error:", err)
	}

	if f != 7 {
		t.Error("乘法錯誤")
	}

}
func TestAdd(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	// mockMale := mock.NewMockMale(ctl)
	mockMale := gotest.NewMockdataInterface(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Add(float64(4), float64(2)).Return(float64(5), nil),
	)
	c := UseData{mockMale}
	f, err := c.UseAdd(3, 2)
	if err != nil {
		t.Error("err, want error:", err)
	}

	if f != 7 {
		t.Error("乘法錯誤")
	}
}
