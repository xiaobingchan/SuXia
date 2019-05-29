package main

import (
	"fmt"
)

const pi = 3.1415

const (
	x = iota // x == 0
	y = iota // y == 1
	z        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

func main() {
	fmt.Printf("%f\n", pi)
	fmt.Printf("%d\n", z)
}
