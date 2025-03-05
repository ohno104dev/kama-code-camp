package day17

// 遞迴解
// 注意題目適合哪種序
// 建立新tree
// 1. 確認函數的參數與返回值
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 2. 確認終止條件
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	// 3. 確認單層遞迴的邏輯(此題適合前 序)
	node := &TreeNode{}
	node.Val = root1.Val + root2.Val
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)

	return node
}

// // BFS解
// // 使用tree1
// func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
//     queue := make([]*TreeNode,0)
//     if root1 == nil{
//         return root2
//     }
//     if root2 == nil{
//         return root1
//     }
//     queue = append(queue,root1)
//     queue = append(queue,root2)

//     for size := len(queue); size>0; size=len(queue) {
//         node1 := queue[0]
//         queue = queue[1:]
//         node2 := queue[0]
//         queue = queue[1:]
//         node1.Val += node2.Val

//         if node1.Left != nil && node2.Left != nil {
//             queue = append(queue,node1.Left, node2.Left)
//         }

//         if node1.Right !=nil && node2.Right !=nil {
//             queue = append(queue, node1.Right, node2.Right))
//         }

//         if node1.Left == nil {
//             node1.Left = node2.Left
//         }

//         if node1.Right == nil {
//             node1.Right = node2.Right
//         }
//     }

//     return root1
// }
