package avl

type AVLNode struct {
	Val    int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

func Add(val int, root *AVLNode) *AVLNode {
	if root == nil {
		return &AVLNode{
			Val: val,
		}
	}

	if val < root.Val {
		root.Left = Add(val, root.Left)
		if height(root.Left)-height(root.Right) == 2 {
			if val < root.Left.Val {
				root = rotateLL(root)
			} else {
				root = rotateLR(root)
			}
		}
	} else {
		root.Right = Add(val, root.Right)
		if height(root.Right)-height(root.Left) == 2 {
			if val > root.Right.Val {
				root = rotateRR(root)
			} else {
				root = rotateRL(root)
			}
		}
	}

	root.Height = max(height(root.Left), height(root.Right)) + 1

	return root
}

func Remove(val int, root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = Remove(val, root.Left)
		if height(root.Right)-height(root.Left) == 2 {
			if height(root.Right.Right) > height(root.Right.Left) {
				root = rotateRR(root)
			} else {
				root = rotateRL(root)
			}
		}
	} else if val > root.Val {
		root.Right = Remove(val, root.Right)
		if height(root.Left)-height(root.Right) == 2 {
			if height(root.Left.Left) > height(root.Left.Right) {
				root = rotateLL(root)
			} else {
				root = rotateLR(root)
			}
		}
	} else {
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			top := root.Right
			for top.Left != nil {
				top = top.Left
			}
			root.Val = top.Val
			root.Right = Remove(top.Val, root)
			if height(root.Left)-height(root.Right) == 2 {
				if height(root.Left.Left) > height(root.Left.Right) {
					root = rotateLL(root)
				} else {
					root = rotateLR(root)
				}
			}
		}
	}

	root.Height = max(height(root.Left), height(root.Right)) + 1

	return root
}

func rotateLL(node *AVLNode) *AVLNode {
	top := node.Left
	node.Left = top.Right
	top.Right = node

	node.Height = max(height(node.Left), height(node.Right)) + 1
	top.Height = max(height(top.Left), height(top.Right)) + 1

	return top
}

func rotateRR(node *AVLNode) *AVLNode {
	top := node.Right
	node.Right = top.Left
	top.Left = node

	node.Height = max(height(node.Left), height(node.Right)) + 1
	top.Height = max(height(top.Left), height(top.Right)) + 1

	return top
}

func rotateLR(node *AVLNode) *AVLNode {
	node.Left = rotateRR(node.Left)
	return rotateLL(node)
}

func rotateRL(node *AVLNode) *AVLNode {
	node.Right = rotateLL(node.Right)
	return rotateRR(node)
}

func height(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
