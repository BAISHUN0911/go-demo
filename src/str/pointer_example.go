package str

import "fmt"

func PointerExample() {
	i := 1
	fmt.Println(&i) // 变量的内存地址s

	p1 := &i
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1 = 2
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(i)

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Println(ptr == nil)

	arr := [3]int{1,2,3}

	for i, v := range &arr { 
		// 1.25版本的go语言这里循环打印&v并不相同，因为每次迭代都会新建新的v变量
		fmt.Println(i, v, &v)
	}

}
