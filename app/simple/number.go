package simple

import "strconv"
import "fmt"


type number struct {
	Value int
}

type numberInterface interface {
	To_s() string
	Inspect() string
	IsReducible() bool
}

func Number(value int) *number {
	number := &number{Value: value}
	return number
} 

func (n *number) To_s() string {
	fmt.Println(strconv.Itoa(n.Value))
	return strconv.Itoa(n.Value)
}

func  (n *number) Inspect() string {
	t := fmt.Sprintf("<< Number >>")
	fmt.Println(t)
	return t
}

func (n number) IsReducible() bool {
	res := false
	fmt.Println(res)
	return res
}