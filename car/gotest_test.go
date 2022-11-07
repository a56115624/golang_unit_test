package car

import (
	gotest "golang_shane_test/car/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

// func Test_Division_1(t *testing.T) {
// 	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
// 		t.Error("除法函式測試沒透過") // 如果不是如預期的那麼就報錯
// 	} else {
// 		t.Log("第一個測試通過了") //記錄一些你期望記錄的資訊
// 	}
// }

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
