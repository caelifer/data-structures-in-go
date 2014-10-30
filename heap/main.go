package main

import (
	"fmt"

	"github.com/caelifer/data-structures-in-go/heap/heap"
)

type IntNode int

// Implement heap.Node interface
func (n IntNode) Less(other heap.Node) bool {
	return n < other.(IntNode)
}

var testInts = [][]IntNode{
	{1},
	{2, 1},
	{5, 15, 95, 40, 21, 1},
}

type RuneNode rune

// Implement heap.Node interface
func (r RuneNode) Less(other heap.Node) bool {
	return r < other.(RuneNode)
}

var testRunes = [][]RuneNode{
	{'a'},
	{'z', 'x', 'y'},
	{'М', 'и', 'р', '!', 'М', 'М', 'М', '世', '界'},
}

type StringNode string

func (s StringNode) Less(other heap.Node) bool {
	return s < other.(StringNode)
}

var testStrings = [][]StringNode{
	{"Hello"},
	{"Hello", "world", "!"},
	{"Здравствуй", "strange", "世界", "!"},
}

func main() {

	// Test ints
	for N, nums := range testInts {
		h := heap.New()
		for _, n := range nums {
			h.Push(n)
		}

		// Check heap
		// fmt.Printf("Heap after test #%d: '%+v'\n", N, h)
		fmt.Printf("Test #%d sorted order: ", N+1)
		for n := h.Pop(); n != nil; n = h.Pop() {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}

	// Test runes
	for N, runes := range testRunes {
		h := heap.New()
		for _, n := range runes {
			h.Push(n)
		}

		// Check heap
		// fmt.Printf("Heap after test #%d: '%+v'\n", N, h)
		fmt.Printf("Test #%d sorted order: ", N+1)
		for n := h.Pop(); n != nil; n = h.Pop() {
			fmt.Printf("%c ", n)
		}
		fmt.Println()
	}

	// Test strings
	for N, strings := range testStrings {
		h := heap.New()
		for _, n := range strings {
			h.Push(n)
		}

		// Check heap
		// fmt.Printf("Heap after test #%d: '%+v'\n", N, h)
		fmt.Printf("Test #%d sorted order: ", N+1)
		for n := h.Pop(); n != nil; n = h.Pop() {
			fmt.Printf("%q ", n)
		}
		fmt.Println()
	}
}
