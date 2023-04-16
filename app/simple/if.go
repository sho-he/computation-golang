package simple

import "fmt"

type _if struct {
	conditon numberInterface
	consequence stateInterface
	alternative stateInterface
	environment environment
}

func If(conditon numberInterface, consequence stateInterface, alternative stateInterface) *_if {
	return &_if{conditon: conditon, consequence: consequence, alternative: alternative}
}

func (i *_if) To_s() string {
	return fmt.Sprintf("if (%v) { %v } else { %v }", i.conditon.To_s(), i.consequence.To_s(), i.alternative.To_s())
}

func (i *_if) Inspect() {
	fmt.Println("<< If >>")
}

func (i *_if) IsReducible() bool {
	return true
}

func (i *_if) Expression(e environment) {
	fmt.Printf("if (%v) { %v } else { %v } %v\n", i.conditon.To_s(), i.consequence.To_s(), i.alternative.To_s(), e)
}

func (i *_if) Reduce(env environment) stateInterface {
	if i.conditon.IsReducible() {
		return If(i.conditon.Reduce(env), i.consequence, i.alternative)
	} else {
		switch i.conditon.(type) {
		case *boolean:
			b := i.conditon.(*boolean)
			if b.Value {
				return i.consequence
			} else {
				return i.alternative
			}
		}
	}
	return nil
}