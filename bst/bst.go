package bst

import (
	"fmt"
)

type BST struct {
	root *Node
	size int
}

func (bst *BST) GetSize() int {
	return bst.size
}

func (bst *BST) AddNonRecursive(val int) {
	if bst.root == nil {
		bst.root = &Node{Val: val}
		bst.size++
		return
	}

	p := bst.root
	for p != nil {
		if val > p.Val {
			if p.Right == nil {
				p.Right = &Node{Val: val}
				bst.size++
				return
			}
			p = p.Right
		} else if val < p.Val {
			if p.Left == nil {
				p.Left = &Node{Val: val}
				bst.size++
				return
			}
			p = p.Left
		} else {
			return
		}
	}
}

func (bst *BST) add(node *Node, val int) *Node {
	if node == nil {
		bst.size++
		return &Node{Val: val}
	}

	if val > node.Val {
		node.Right = bst.add(node.Right, val)
	} else if val < node.Val {
		node.Left = bst.add(node.Left, val)
	}
	return node
}

func (bst *BST) Add(val int) {
	bst.root = bst.add(bst.root, val)
}

func (bst *BST) contains(node *Node, val int) bool {
	if node == nil {
		return false
	}
	if val == node.Val {
		return true
	} else if val < node.Val {
		return bst.contains(node.Left, val)
	} else {
		return bst.contains(node.Right, val)
	}
}

func (bst *BST) Contains(val int) bool {
	return bst.contains(bst.root, val)
}

func (bst *BST) PreOrderNonRecursive() {
	if bst.root == nil {
		return
	}

	var stack []*Node
	stack = append(stack, bst.root)

	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]

		fmt.Println(node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
}

func (bst *BST) preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	bst.preOrder(node.Left)
	bst.preOrder(node.Right)
}

func (bst *BST) PreOrder() {
	bst.preOrder(bst.root)
}

func (bst *BST) inOrder(node *Node) {
	if node == nil {
		return
	}

	bst.inOrder(node.Left)
	fmt.Println(node.Val)
	bst.inOrder(node.Right)
}

func (bst *BST) InOrder() {
	bst.inOrder(bst.root)
}

func (bst *BST) postOrder(node *Node) {
	if node == nil {
		return
	}

	bst.postOrder(node.Left)
	bst.postOrder(node.Right)
	fmt.Println(node.Val)
}

func (bst *BST) PostOrder() {
	bst.postOrder(bst.root)
}

func (bst *BST) LevelOrder() {
	if bst.root == nil {
		return
	}

	var q []*Node
	q = append(q, bst.root)

	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		fmt.Println(node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
}

func (bst *BST) minimum(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return bst.minimum(node.Left)
}

func (bst *BST) Minimum() int {
	if bst.root == nil {
		panic("BST is empty")
	}
	return bst.minimum(bst.root).Val
}

func (bst *BST) maximum(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return bst.maximum(node.Right)
}

func (bst *BST) Maximum() int {
	if bst.root == nil {
		panic("BST is empty")
	}
	return bst.maximum(bst.root).Val
}

func (bst *BST) removeMin(node *Node) *Node {
	if node.Left == nil {
		rightNode := node.Right
		node.Right = nil
		bst.size--
		return rightNode
	}
	node.Left = bst.removeMin(node.Left)
	return node
}

func (bst *BST) RemoveMin() int {
	ret := bst.Minimum()
	bst.root = bst.removeMin(bst.root)
	return ret
}

func (bst *BST) removeMax(node *Node) *Node {
	if node.Right == nil {
		leftNode := node.Left
		node.Right = nil
		bst.size--
		return leftNode
	}
	node.Right = bst.removeMin(node.Right)
	return node
}

func (bst *BST) RemoveMax() int {
	ret := bst.Maximum()
	bst.root = bst.removeMax(bst.root)
	return ret
}

func (bst *BST) remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if val < node.Val {
		node.Left = bst.remove(node.Left, val)
		return node
	} else if val > node.Val {
		node.Right = bst.remove(node.Right, val)
		return node
	} else {
		if node.Left == nil {
			rightNode := node.Right
			node.Right = nil
			bst.size--
			return rightNode
		}
		if node.Right == nil {
			leftNode := node.Left
			node.Left = nil
			bst.size--
			return leftNode
		}

		successor := bst.minimum(node.Right)
		successor.Right = bst.removeMin(node.Right)
		successor.Left = node.Left
		node.Left = nil
		node.Right = nil

		return successor
	}
}

func (bst *BST) Remove(val int) {
	bst.root = bst.remove(bst.root, val)
}