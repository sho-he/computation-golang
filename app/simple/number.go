package simple

import "strconv"
import "fmt"


type number struct {
	Value int
}

func Number(value int) *number {
	number := &number{Value: value}
	return number
} 

func (n *number) To_s() string {
	return strconv.Itoa(n.Value)
}

func  (n *number) Inspect() {
	t := fmt.Sprintf("<< Number >>")
	fmt.Println(t)
}

func (n *number) Expression() {
	fmt.Println(n.To_s())
}

func (n number) IsReducible() bool {
	return false
}

func (n *number) Reduce() numberInterface {
	return n
}

func (n *number) getNumber() int {
	return n.Value
}