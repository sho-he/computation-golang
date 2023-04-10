package simple

import "strconv"
import "fmt"


type Number struct {
	Value int
}

func NewNumber(value int) *Number {
	number := &Number{Value: value}
	return number
} 

func (n *Number) To_s() string {
	fmt.Println(strconv.Itoa(n.Value))
	return strconv.Itoa(n.Value)
}

func  (n *Number) Inspect() string {
	t := fmt.Sprintf("<< %T >>", n)
	fmt.Println(t)
	return t
}
