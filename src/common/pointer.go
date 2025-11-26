package common

import "fmt"

func PointerDemo() {
	i := 100
	fmt.Println("局部变量i的内存地址= ", &i)
	i2 := i
	var i3 int = i // 值拷贝:值一样,内存地址一定不同
	var p1 *int = &i
	p2 := p1
	fmt.Println("局部变量i6的内存地址= ", &i3, "i6的值= ", i3)
	fmt.Println("局部变量i5的内存地址= ", &i2, "i5的值= ", i2)
	fmt.Println("指针p1的内存地址= ", &p1, "p1指向的值= ", *p1, "p1=", p1)
	fmt.Println("指针p2的内存地址= ", &p2, "p2指向的值= ", *p2, "p2=", p2)
}
