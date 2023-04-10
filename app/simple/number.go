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

func  (n Number) Inspect() string {
	t := fmt.Sprintf("<< %T >>", n)
	fmt.Println(t)
	return t
}
