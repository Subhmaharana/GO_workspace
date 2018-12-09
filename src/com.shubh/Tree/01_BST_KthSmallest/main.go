package main

type treeNode struct {
	value int
	left  treeNode
	right treeNode
}

func treeNode_new(val int) treeNode {
	var node treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

type myStack []*treeNode

func (s myStack) Push(v *treeNode) myStack {
	return append(s, v)
}

func (s myStack) Pop() (myStack, *treeNode) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

func main() {

	kthSmallest()
}

func kthSmallest(A *treeNode, K int) int {
	curr := A

	stk := myStack{}

	for curr != nil {
		stk = stk.Push(curr)
		curr = curr.left
	}

	c := 0
	for !stk.isEmpty() {
		stk, node := stk.Pop()
		c++
		if c == K {
			return node.value
		}

		if node.right != nil {
			node = node.right
			for node != nil {
				stk = stk.Push(node)
				node = node.left
			}
		}
	}
	return -1

}
