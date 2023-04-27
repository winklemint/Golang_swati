package main

import "fmt"

type calculator interface {
	Addition()
	Subtraction()
	Multiplication()
	Division()
}
type Data struct {
	Num1 int
	Num2 int
}
type Data1 struct {
	Num1 float64
	Num2 float64
}

func (a Data) Addition() int {
	return a.Num1 + a.Num2
}
func (s Data) Subtraction() int {
	return s.Num1 - s.Num2
}
func (m Data) Multiplication() int {
	return m.Num1 * m.Num2
}
func (d Data) Division() int {
	return d.Num1 / d.Num2
}

func (a Data1) Addition() float64 {
	return a.Num1 + a.Num2
}
func (s Data1) Subtraction() float64 {
	return s.Num1 - s.Num2
}
func (m Data1) Multiplication() float64 {
	return m.Num1 * m.Num2
}
func (d Data1) Division() float64 {
	return d.Num1 / d.Num2
}

func Result(c Data) {
	fmt.Println(c.Addition())
	fmt.Println(c.Subtraction())
	fmt.Println(c.Multiplication())
	fmt.Println(c.Division())
}
func Result1(c Data1) {
	fmt.Println(c.Addition())
	fmt.Println(c.Subtraction())
	fmt.Println(c.Multiplication())
	fmt.Println(c.Division())
}
func main() {
	a := Data{Num1: 2, Num2: 6}
	b := Data1{Num1: 2, Num2: 6}
	Result(a)
	Result1(b)
}
