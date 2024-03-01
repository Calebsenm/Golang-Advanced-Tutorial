// Go does not support method Overloading

package main

type maths struct{}

func (m *maths) add(a, b int) int {
	return a + b
}

func (m *maths) add(a, b, c int) int {
	return a + b + c
}

func main() {
	m := &maths{}
}
