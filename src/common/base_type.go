package common

import "fmt"

// 定义一个常量
const CACHE_SIZE = 1000

func BaseTypeDemo() {
	var b bool = true
	// 同一类型的多个变量可以声明在同一行
	var i, i2 int = 1, 2
	var i8 int8 = 127   // 有符号 8 位整型 (-128 到 127)
	var ui8 uint8 = 255 // 无符号 8 位整型 (0 到 255)
	var f1 float32 = 3.14
	var s1 string = "hello world"

	// 方法一：使用fmt.Sprintf格式化
	result := fmt.Sprintf("%s %t %d %d %d %d %.2f", s1, b, i, i2, i8, ui8, f1)

	// 所有声明的局部变量必须被使用，否则编译报错
	// 接收但故意丢弃这个值
	_ = result
	// 只获取函数返回值的后两个
	_, f2, s2 := numbers()
	fmt.Println("numbers()函数后两个返回值: ", f2, s2)

	// 没有初始化就为零值
	var i3 int
	fmt.Println("int的默认值是: ", i3)

	var b2 bool
	fmt.Println("bool的默认值是: ", b2) // bool 类型的零值为 false

	// 自动判定变量类型
	var v_name = true
	fmt.Println(v_name)

	// 变量只声明一次
	// i4 := 1 相当于var i4 int; i4 = 1
	i4 := 1
	_ = i4

}

//一个可以返回多个值的函数
func numbers() (int, float64, string) {
	a, b, c := 1, 2.0, "str"
	return a, b, c
}
