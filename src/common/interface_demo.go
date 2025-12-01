package common

import "fmt"

func InterfaceDemo() {
	var a Animal = Dog{}

	fmt.Println(a)

	// Cat结构体没有实现Animal接口要求的Speak()方法，因此无法赋值给Animal接口类型变量。
	// var b Animal = Cat{}
	// fmt.Println(b)
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (Dog) Speak() string {
	return "汪汪!"
}

type Cat struct{}

func (Cat) Color() string {
	return "我是白色的"
}
