package common

import (
	"fmt"
	"strings"
)

func SliceDemo() {
	// This is a demo function for slice operations in Go.
	var s1 []int
	s1 = append(s1, 1, 3, 2)

	for i, v := range s1 {
		fmt.Println("i的地址: ", &i, " i的值: ", i)
		fmt.Println("v的地址: ", &v, " v的值: ", v, "原 slice 元素的地址", &s1[i])
	}

	// 对int进行变量
	for i := range 3 {
		fmt.Println("i的地址: ", &i, " i的值: ", i)
	}

}

func SliceDemo2() {
	keyStr := "iCoach:track:2025-12-24:course:count:0:999"
	parts := strings.Split(keyStr, ":")

	fmt.Println(parts)
	fmt.Println(parts[4])
	fmt.Println(parts[7-1])

}
