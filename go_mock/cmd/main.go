package main

import (
	"fmt"
	"prac/go_mock/calc"
)

func main() {
	opt := calc.NewSelectedOption(calc.SelectedOption{})
	fmt.Println(opt.Sum(123, 98))
}
