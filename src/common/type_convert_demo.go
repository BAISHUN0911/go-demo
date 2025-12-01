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
