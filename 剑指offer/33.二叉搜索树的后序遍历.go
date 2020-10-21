package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	root := len(postorder)-1
	for i := 0; i < len(postorder); i++ {
		if postorder[i] >= postorder[len(postorder)-1] {
			root = i
		}
		if i > root && postorder[i] < postorder[len(postorder)-1] {
			return false
		}
	}
	return verifyPostorder(postorder[:root]) && verifyPostorder(postorder[root:len(postorder)-1])
}


