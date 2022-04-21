package calc

type Calculate interface {
	Sum(first, second int) int
	Delta(first, second int) int
}

type SelectedOption struct {
	Calc Calculate
}

func (s SelectedOption) Sum(first, second int) int {
	return first + second
}

func (s SelectedOption) Delta(first, second int) int {
	return first - second
}

func NewSelectedOption(calc Calculate) *SelectedOption {
	return &SelectedOption{Calc: calc}
}
