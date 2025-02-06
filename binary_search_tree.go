package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) bool {
	if value < 0 {
		return false
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}

	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}

	return true
}

func (n *Node) Search(value int) bool {
	if value < 0 {
		return false
	}

	if value == n.Value {
		return true
	}
	if value < n.Value {
		return n.Left.Search(value)
	}
	if value > n.Value {
		return n.Right.Search(value)
	}

	return false
}

func (n *Node) printOdered() {
	if n == nil {
		return
	}

	n.Left.printOdered()
	fmt.Print(n.Value, "->")
	n.Right.printOdered()
}

func main() {
	root := Node{Value: 10}

	values := []int{100, 20, 70, 40, 0, -1, 50, 5}
	for _, value := range values {
		root.Insert(value)
	}

	root.printOdered()

	fmt.Println("\n ================================== \n")

	exist := root.Search(5)
	fmt.Printf("expected(true)->got(%t)", exist)

}
