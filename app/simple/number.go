package simple

import "strconv"
import "fmt"


type number struct {
	Value int
}

func Number(value int) number {
	number := number{Value: value}
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
