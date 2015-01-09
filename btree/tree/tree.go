package tree

type Node interface {
	Less(other Node) bool
}

type CustomFn func(Node)

type Interface interface {
	Insert(Node)
	Walk(CustomFn)
}

type binaryTree struct {
	node  Node
	left  Interface
	right Interface
}

func New() Interface {
	return &binaryTree{}
}

func (t *binaryTree) Insert(node Node) {
	if t.node == nil {
		t.node = node
		t.right = New()
		t.left = New()
	} else {
		if node.Less(t.node) {
			t.left.Insert(node)
		} else {
			t.right.Insert(node)
		}
	}
}

func (t *binaryTree) Walk(fn CustomFn) {
	if t.node == nil {
		return
	}

	t.left.Walk(fn)
	fn(t.node)
	t.right.Walk(fn)
}
