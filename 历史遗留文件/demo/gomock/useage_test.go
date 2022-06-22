package main

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	// 断言Bar()被执行
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// 断言99被传入Bar()
	// Anything else will fail.
	// 其他情况将报出异常
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(102)

	SUT(m)
}
