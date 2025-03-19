package day31

// Time Complexity: O(N)
// Space Complexity: O(N)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Status int

const (
	NO_COVER Status = iota
	HAS_CAMERA
	COVERED
)

func minCameraCover(root *TreeNode) int {
	res := 0
	if traversal(root, &res) == NO_COVER {
		res++
	}

	return res
}

func traversal(cur *TreeNode, res *int) Status {
	// 當空節點為有覆蓋的情況才能得到最少的結果
	if cur == nil {
		return COVERED
	}

	left := traversal(cur.Left, res)
	right := traversal(cur.Right, res)

	if left == COVERED && right == COVERED {
		return NO_COVER
	}

	if left == NO_COVER || right == NO_COVER {
		*res++
		return HAS_CAMERA
	}

	if left == HAS_CAMERA || right == HAS_CAMERA {
		return COVERED
	}

	return -1
}
