package main_test

import (
	"fmt"
	"testing"

	nodes "./core/nodes"
)

func testNode1() bool {
	b1, b2 := nodes.NewNodeBody(), nodes.NewNodeBody()
	ln1 := nodes.NewNodeLine()
	b1.InsertPoint(nodes.NODE_TYPE_OUTPUT, "o1")
	for i := 0; i < 100000; i++ {
		b2.InsertPoint(nodes.NODE_TYPE_INPUT, fmt.Sprintf("i%d", i))
		b2.InsertPoint(nodes.NODE_TYPE_OUTPUT, fmt.Sprintf("o%d", i))
	}
	temp1, ok := b1.GetPoint(nodes.NODE_TYPE_OUTPUT, "o1")
	if !ok {
		return false
	}
	temp2, ok := b2.GetPoint(nodes.NODE_TYPE_INPUT, "i600")
	if !ok {
		return false
	}
	ln1.Connect(temp1, temp2)
	return true
}
func TestNode1(t *testing.T) {

	if r := testNode1(); !r {
		t.FailNow()
	}
}

func BenchmarkNode1(b *testing.B) {
	if r := testNode1(); !r {
		b.FailNow()
	}
}
