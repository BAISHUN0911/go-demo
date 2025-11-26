// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

//  import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import (
	"fmt"
	"go-demo/src/common"
)

/*
func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数
（如果有 init() 函数则会先执行该函数）
*/
func main() {
	fmt.Println("Hello, Go on Windows!")
	common.Hello()
}
