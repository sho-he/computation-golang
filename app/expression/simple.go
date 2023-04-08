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

func  (n *Number) Inspect() {
	fmt.Printf("<<# {%T} >>", n.self)
}


type Add struct {
	Left int
	Right int
}

func (a *Add) To_s() string {
	return fmt.Sprintf("%d + %d", a.Left, a.Right)
}

func (a *Add) Inspect() {
	fmt.Printf("<<# {%T} >>", a)
}


type Multiply struct {
	Left int
	Right int
}

func (m *Multiply) to_s() string {
	return fmt.Sprintf("%d * %d", m.Left, m.Right)
}

func (m *Multiply) Inspect() {
	fmt.Printf("<<# {%T} >>", m)
}