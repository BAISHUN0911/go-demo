package str

import (
	"encoding/base64"
	"fmt"
)

func FormatStr() {
	str := "5Yqg5YWlR1ZB5Lqk5rWB576k"
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(decodeBytes, err)
	// 打印结果
	// %s 会将字节切片以字符串形式输出
	// %v 会以默认格式输出错误信息，如果没有错误，会显示 <nil>
	fmt.Printf("解码后的字节切片 (bytes): %v\n", decodeBytes)
	fmt.Printf("解码后的字符串: %s\n", decodeBytes)
	fmt.Printf("错误信息: %v\n", err)
	fmt.Printf("错误信息: %s\n", err)

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	// 格式化参数生成格式化的字符串并写入标准输出
	fmt.Printf(url, stockcode, enddate)
}
