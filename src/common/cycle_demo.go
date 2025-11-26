package common

import "fmt"

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
