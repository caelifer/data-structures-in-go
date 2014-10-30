package main

import "data-structures/btree/tree"

type IntNode int

func (n IntNode) Less(other tree.Node) bool {
	// TODO: need to handle type assertion error
	return n < other.(IntNode)
}

type StringNode string

func (n StringNode) Less(other tree.Node) bool {
	// TODO: need to handle type assertion error
	return n < other.(StringNode)
}

func main() {
	intBT := tree.New()
	intBT.Insert(IntNode(5))
	intBT.Insert(IntNode(1))
	intBT.Insert(IntNode(7))
	intBT.Insert(IntNode(2))
	intBT.Insert(IntNode(4))
	intBT.Insert(IntNode(6))
	intBT.Insert(IntNode(3))
	intBT.Walk()

	stringBT := tree.New()
	stringBT.Insert(StringNode("ello"))
	stringBT.Insert(StringNode("der"))
	stringBT.Insert(StringNode("ost"))
	stringBT.Insert(StringNode("a"))
	stringBT.Insert(StringNode("cy"))
	stringBT.Insert(StringNode("z"))
	stringBT.Insert(StringNode("j"))
	stringBT.Walk()
}
