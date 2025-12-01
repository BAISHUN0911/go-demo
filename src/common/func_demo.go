package common

import (
	"errors"
	"fmt"
	"math"
)

func FuncDemo() {
	// 函数支持作为变量
	var f func(int, int) int // 等同于 f := func(a, b int) int { return a + b }
	f = func(a, b int) int { return a + b }
	fmt.Println(f(1, 2))

	f2 := makeSqrtWithOffset(10)
	fmt.Println(f2(16)) // sqrt(16) + 10
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func DemoFunc() (res int, err error) {
	res, err = div(10, 2)
	if err != nil {
		// 处理错误
		return
	}
	// 继续处理 res,继续业务逻辑
	return
}

// 闭包示例
func makeSqrtWithOffset(offset float64) func(float64) float64 {
	return func(x float64) float64 {
		return math.Sqrt(x) + offset
	}
}
