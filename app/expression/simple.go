package simple

import "strconv"
import "fmt"


type Number struct {
	self *Number
	Value int
}

func (n *Number) To_s() string {
	return strconv.Itoa(n.Value)
}

func  (n *Number) Inspect() string {
	t := fmt.Sprintf("<< %T >>", n)
	fmt.Println(t)
	return t
}


type Add struct {
	Left Number
	Right Number
}

func (a *Add) To_s() string {
	expression := fmt.Sprintf("%d + %d", a.Left.Value, a.Right.Value)
	fmt.Println(expression)
	return expression
}

func (a *Add) Inspect() string {
	t := fmt.Sprintf("<< %T >>", a)
	fmt.Println(t)
	return t
}


type Multiply struct {
	Left Number
	Right Number
}

func (m *Multiply) to_s() string {
	expression := fmt.Sprintf("%d * %d", m.Left.Value, m.Right.Value)
	fmt.Println(expression)
	return expression
}

func (m *Multiply) Inspect() string {
	t := fmt.Sprintf("<< %T >>", m)
	fmt.Println(t)
	return t
}