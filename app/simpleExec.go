package main

import "fmt"
import "golang_computation/simple"

func main () {
		result := simple.Add{Left: simple.Number{Value: 1}, Right: simple.Number{Value: 2}}
		fmt.Println(result)
		// result2 := simple.Add{Left: simple.Number{Value: 3}, Right: simple.Number{Value: 6}}
		// result3 := simple.Multiply{Left: result, Right: result2}
		// result3.To_s()
		// result.Inspect()
}
