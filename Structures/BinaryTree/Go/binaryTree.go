package binarytree

type Node struct {
	parent      *Node
	left, right *Node
	Value       int
}

func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) IsParent() bool {
	return n.parent == nil
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SubSearch(value int) *Node {
	return search(value, n)
}

type BT struct {
	root *Node
	size int
}

func NewBT() *BT {
	return new(BT)
}

func (b *BT) Insert(value int) {
	if b.root == nil {
		n := new(Node)
		n.Value = value
		b.root = n
		return
	}
	insert(value, b.root)
}

func (b *BT) Search(key int) *Node {
	return search(key, b.root)
}

func insert(value int, parent *Node) {
	if value < parent.Value {
		if parent.Left() != nil {
			insert(value, parent.Left())
			return
		}
		n := new(Node)
		n.parent = parent
		n.Value = value
		parent.left = n
	} else {
		if parent.Right() != nil {
			insert(value, parent.Right())
			return
		}
		n := new(Node)
		n.parent = parent
		n.Value = value
		parent.right = n
	}
}

func search(key int, node *Node) *Node {
	if node == nil || node.Value == key {
		return node
	}
	if key < node.Value {
		return search(key, node.Left())
	} else {
		return search(key, node.Right())
	}
}
