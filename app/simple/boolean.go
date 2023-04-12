package simple

import "fmt"
import "strconv"

type boolean struct {
	Value bool
}

func Boolean(value bool) *boolean {
	boolean := &boolean{Value: value}
	return boolean
}

func (b *boolean) To_s() string {
	return strconv.FormatBool(b.Value)
}

func (b *boolean) Inspect() {
	t := fmt.Sprintf("<< Boolean >>")
	fmt.Println(t)
}

func (b *boolean) Expression() {
	fmt.Println(b.To_s())
}

func (b *boolean) IsReducible() bool {
	return false
}

func (b *boolean) Reduce() numberInterface {
	return Number(b.getNumber())
}

func (b *boolean) getNumber() int {
	if b.Value == true {
		return 1
	} else {
		return 0
	}
}