package common

import "fmt"

func PointerDemo() {
	i := 100
	fmt.Println("局部变量i的内存地址= ", &i)
	i2 := i
	fmt.Println("局部变量i2的内存地址= ", &i2, "i2的值= ", i2)

	var i3 int = i // 值拷贝:值一样,内存地址一定不同
	fmt.Println("局部变量i3的内存地址= ", &i3, "i3的值= ", i3)

	var p1 *int = &i
	p2 := p1
	fmt.Println("指针p1的内存地址= ", &p1, "p1指向的值= ", *p1, "p1=", p1)
	fmt.Println("指针p2的内存地址= ", &p2, "p2指向的值= ", *p2, "p2=", p2)

	// 指向p1的指针
	p3 := &p1
	fmt.Println("指针p3的内存地址= ", &p3, "p3指向的值= ", **p3, "p3=", p3, "(指针p1的内存地址)")

	// 指针作为函数参数
	i4 := 100
	i5 := 200
	fmt.Printf("交换前 i4 的值 : %d i4的内存地址%x\n", i4, &i4)
	fmt.Printf("交换前 i5 的值 : %d i5的内存地址%x\n", i5, &i5)
	swap(&i4, &i5)
	// 交换后变量的内存地址不变，值交换了
	fmt.Printf("交换后 i4 的值 : %d i4的内存地址%x\n", i4, &i4)
	fmt.Printf("交换后 i5 的值 : %d i5的内存地址%x\n", i5, &i5)

	swapValue(i4, i5)
	fmt.Printf("swapValue后 i4 的值 : %d i4的内存地址%x\n", i4, &i4)
	fmt.Printf("swapValue后 i5 的值 : %d i5的内存地址%x\n", i5, &i5)

	i4, i5 = swapValueV2(i4, i5)
	fmt.Printf("swapValueV2后 i4 的值 : %d i4的内存地址%x\n", i4, &i4)
	fmt.Printf("swapValueV2后 i5 的值 : %d i5的内存地址%x\n", i5, &i5)

}

// 改的是内存内容，地址不变
func swap(x, y *int) {
	// Go 支持多重赋值
	*x, *y = *y, *x
}

// 传的是「值的副本」，函数里交换的只是局部拷贝，外面的 i4 和 i5 完全没被修改
/*
进入 swapValue 函数时
i4(100) → A
i5(200) → B

x(100)  → C
y(200)  → D
x 和 y 是全新的变量，不是 i4、i5 的地址
*/
func swapValue(x, y int) {
	// 只影响函数内部变量
	x, y = y, x
}

func swapValueV2(x, y int) (int, int) {
	return y, x
}
