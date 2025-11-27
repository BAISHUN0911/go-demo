package common

import "fmt"

func FuncDemo() {
	// 函数支持作为变量
	var f func(int, int) int  // 等同于 f := add
	f = add
	fmt.Println(f(1, 2))

}

// 定义 add 函数
func add(a, b int) int {
	return a + b
}
