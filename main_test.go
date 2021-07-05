package main_test

import (
	"testing"

	nodes "./core/nodes"
)

func testNode1() bool {
	p1, p2, p3 := nodes.NodePoint{}, nodes.NodePoint{}, nodes.NodePoint{}
	b1, b2 := nodes.NewNodeBody(), nodes.NewNodeBody()
	ln1 := nodes.NodeLine{}
	b1.InsertPoint(nodes.NODE_TYPE_OUTPUT, "o1", &p1)
	b2.InsertPoint(nodes.NODE_TYPE_INPUT, "i1", &p2)
	b2.InsertPoint(nodes.NODE_TYPE_OUTPUT, "o1", &p3)
	temp1, ok := b1.GetPoint(nodes.NODE_TYPE_OUTPUT, "o1")
	if !ok {
		return false
	}
	temp2, ok := b2.GetPoint(nodes.NODE_TYPE_INPUT, "i1")
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
