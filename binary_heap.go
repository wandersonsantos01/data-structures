package main

import (
	"container/heap"
	"fmt"
)

/*
Problem: Merge K Sorted Arrays (or Lists)
Problem Statement:
Given k sorted arrays, merge them into a single sorted array as efficiently as possible.

Example:
Input:
[
  [1, 4, 7],
  [2, 5, 8],
  [3, 6, 9]
]

Output:
[1, 2, 3, 4, 5, 6, 7, 8, 9]


#########################################

input [][]

output[] -> ordered

input is empty - []
    expect([]) -> got([])

consider negative values
    no - only positives

approach
    A - binary heap
    B - lists (slowest)

#1 - receive input
#2 - validate input
#2.1 - check if it is empty
#2.2 - check if it is negative
#3 - start a heap (native)
#4 - add node
#5 - how to order this
#6 - consider remove node (if we have time)
*/

type Element struct {
	Value      int
	ArrayIdx   int
	ElementIdx int
}

type PriorityQueue []*Element

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(element any) {
	*p = append(*p, element.(*Element))
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func Process(input [][]int) []int {
	output := []int{}

	if len(input) < 1 {
		return output
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for idx, list := range input {
		if len(list) > 0 {
			element := &Element{
				Value:      list[0],
				ArrayIdx:   idx,
				ElementIdx: 0,
			}
			heap.Push(pq, element)
		}
	}

	for pq.Len() > 0 {
		minElem := heap.Pop(pq).(*Element)
		output = append(output, minElem.Value)

		nextIdx := minElem.ElementIdx + 1
		if nextIdx < len(input[minElem.ArrayIdx]) {
			newElement := &Element{
				Value:      input[minElem.ArrayIdx][nextIdx],
				ArrayIdx:   minElem.ArrayIdx,
				ElementIdx: nextIdx,
			}
			heap.Push(pq, newElement)
		}

	}

	return output
}

func main() {
	input := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	fmt.Println(Process(input))
}
