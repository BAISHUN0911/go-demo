package common

import "go-demo/src/common/model"

func OperatorDemo() {
	s1 := "abc"
	s2 := "abc"
	println(s1 == s2) // true,Go 的 string 是 值类型，用 == 就是内容比较

	u1 := model.User{Username: "Alice", Role: "admin"}
	u2 := model.User{Username: "Alice", Role: "admin"}
	println(u1 == u2) // true,如果 struct 包含 slice、map、function 等不可比较字段，则不能直接用 ==
}
