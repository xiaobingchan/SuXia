package main

import (
	"fmt"
)

//示例代码
//var frenchHello string // 声明变量为字符串的一般方法
//var emptyString string = "" // 声明了一个字符串变量，初始化为空字符串

func main5() {
	//	no, yes, maybe := "no", "yes", "maybe" // 简短声明，同时声明多个变量
	japaneseHello := "Konichiwa" // 同上
	//frenchHello = "Bonjour" // 常规赋值
	fmt.Printf("%s\n", japaneseHello) // %s代表字符串

	s := "hello"
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
}
