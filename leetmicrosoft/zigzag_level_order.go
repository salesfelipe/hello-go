package leetmicrosoft

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
}

func generateTree(input []int) *TreeNode {
	var root *TreeNode = &TreeNode{Val: input[0]}

	var prev *TreeNode = root
	for i := 1; i < len(input); i++ {
		node := &TreeNode{Val: input[i]}

		if i%2 == 0 {
			prev.Right = node
		} else {
			prev.Left = node
		}

		leftIndex := 2*i + 1

		// insert node
		current.Left = node
	}

	return root
}
