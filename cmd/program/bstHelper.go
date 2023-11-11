package program

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func buildTree(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Value: preorder[0]}
	i := 1
	for ; i < len(preorder); i++ {
		if preorder[i] > root.Value {
			break
		}
	}

	root.Left = buildTree(preorder[1:i])
	root.Right = buildTree(preorder[i:])

	return root
}

func printTree(root *TreeNode, prefix string, isLast bool, position string) {
	if root != nil {
		fmt.Print(prefix)
		if isLast {
			fmt.Printf("└── %s: %d\n", position, root.Value)
			prefix += "    "
		} else {
			fmt.Printf("├── %s: %d\n", position, root.Value)
			prefix += "│   "
		}

		printTree(root.Left, prefix, root.Right == nil, "L")
		printTree(root.Right, prefix, true, "R")
	}
}

func PrettyPrintTree(preorder []int) {
	root := buildTree(preorder)
	printTree(root, "", true, "root")
}
