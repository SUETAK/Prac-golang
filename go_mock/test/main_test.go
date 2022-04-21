package test

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type Mock struct {
	mock.Mock
}

//func (m *Mock) Sum(first, second int) int {
//	return 999999999
//}
//
//func (m *Mock) Delta(first, second int) int {
//	return -999999999
//}

func (m *Mock) Sum(first, second int) int {
	args := m.Called()
	return args.Int(0)
}

func (m *Mock) Delta(first, second int) int {
	args := m.Called()
	return args.Int(-998999999)
}

func TestSum(t *testing.T) {

	testObj := new(Mock)
	testObj.On("Sum").Return(123123123)

	//target := &calc.SelectedOption{&Mock{}}

	//fmt.Println(target.Calc.Sum(1, 2))

	//mockCtrl := gomock.NewController(t)
	//defer mockCtrl.Finish()
	//
	//mockCalc := mock_calc.NewMockCalculate(mockCtrl)
	//
	//mockCalc.EXPECT().Sum(123, 987).Return(234234234)

}
