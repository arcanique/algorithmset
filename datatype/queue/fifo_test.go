package queue

import (
	"fmt"
	"testing"
)

func TestQueue_NewFromSlice(t *testing.T) {
	var x []interface{}

	a := []int{1,2,3}
	x = make([]interface{},len(a))
	for i, e := range a {
		x[i] = e
	}
	f,_ := NewFromSlice(x)

	fmt.Printf("%v", f)
}