package main

import (
	"fmt"
	"prac/go_mock/calc"
)

func main() {
	opt := calc.NewSelectedOption()
	fmt.Println(opt.Sum(123, 98))
}
