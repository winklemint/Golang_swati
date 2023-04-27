package calci

type Calc struct {
	Num1 int
	Num2 int
}

func (c *Calc) Add() int {
	return c.Num1 + c.Num2
}

func (c *Calc) Sub() int {
	return c.Num1 - c.Num2
}
func (c *Calc) Multi() int {
	return c.Num1 * c.Num2
}
func (c *Calc) Div() int {
	return c.Num1 / c.Num2
}
