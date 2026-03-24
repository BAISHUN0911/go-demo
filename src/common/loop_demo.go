package common

import "fmt"

// 使用range遍历数组时，遍历的每项元素内存地址是不同的
func AddressWithinRange() {
	// 1. 在函数内部声明结构体
	type Message struct {
		ID      int
		Content string
	}
	// 2. 创建结构体切片（通常比固定长度数组更常用）
	var msgs []Message
	// 3. 放入元素
	// 方式 A：使用字面量直接添加
	msgs = append(msgs, Message{ID: 1, Content: "A"})
	// 4. 另一种初始化方式：带初始值的声明
	moreMsgs := []Message{
		{ID: 2, Content: "B"},
		{ID: 3, Content: "C"},
	}
	// 合并
	msgs = append(msgs, moreMsgs...)

	for index, value := range msgs {
		fmt.Println("index:", index, "value:", value)
		fmt.Println("i地址", &index, "value地址", &value)
	}
}

func CycleDemo() {
	// 数组默认值0
	var numbers0 [5]int
	fmt.Println(numbers0[0])

	numbers := [5]int{1, 2, 5, 4, 3}
	// for循环打印
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}
	// range方式
	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}

	// 冒泡排序
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println(numbers)

	// 二维数组
	values := [][]int{}
	// 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
	fmt.Println(values[1][1])

	// range循环二维数组
	for i, row := range values {
		for j, value := range row {
			fmt.Printf("Row %d Column %d Value %d\n", i, j, value)
		}
	}

}
