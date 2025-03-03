package day15

// 遞迴解
// 注意題目適合哪種序
func countNodes(root *TreeNode) int {
	return traversal(root)
}

// // 一般二叉樹迭代解
// func countNodes(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }

//     stack := []*TreeNode{root}
//     count := 0
//     for len(stack) > 0 {
//         node := stack[len(stack) - 1]
//         stack = stack[:len(stack) - 1]
//         if node.Left != nil {
//             stack = append(stack, node.Left)
//         }
//         if node.Right != nil {
//             stack = append(stack, node.Right)
//         }
//         count++
//     }

//     return count
// }

// // 一般二叉樹解
// // 1. 確認函數的參數與返回值
// func traversal(node *TreeNode) int {
// 	// 2. 確認終止條件
// 	if node == nil {
// 		return 0
// 	}

// 	// 3. 確認單層遞迴的邏輯(此題適合後 序)
// 	l := traversal(node.Left)
// 	r := traversal(node.Right)

// 	return l + r + 1
// }

// 完全二叉樹特性解
// Time Complexity: O(Log(N))
// 1. 確認函數的參數與返回值
func traversal(node *TreeNode) int {
	// 2. 確認終止條件
	if node == nil {
		return 0
	}
	left := node.Left
	right := node.Right
	ldep, rdep := 0, 0
	for left != nil {
		left = left.Left
		ldep++
	}
	for right != nil {
		right = right.Right
		rdep++
	}
	if ldep == rdep {
		return 2<<ldep - 1
	}

	// 3. 確認單層遞迴的邏輯(此題適合後 序)
	l := traversal(node.Left)
	r := traversal(node.Right)

	return l + r + 1
}
