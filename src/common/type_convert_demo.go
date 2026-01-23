package common

import "fmt"

func TypeConvertDemo() {
	var i interface{} = "Hello, World"
	// 类型断言用于将接口类型转换为指定类型
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}

func TypeJudge() {
	var id interface{} = 123
	var idValue uint
	// type 关键字只用于匹配分支，不参与运行时计算
	switch t := id.(type) {
	// 当 id 的动态类型是 int 时，进入该分支，并把值赋给 t
	case int:
		idValue = uint(t)
	}
	fmt.Println(idValue)
}
