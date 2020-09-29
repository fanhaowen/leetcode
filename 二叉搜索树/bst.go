package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	data int
	LeftTree *TreeNode
	RightTree *TreeNode
}
var pre *TreeNode
func main() {
	head := TreeNode{4,nil,nil}
	TreeInit(&head)
	root := treeToDoublyList(&head)
	//tail := root.LeftTree
	for i := 0; i <=4 ; i++ {
		fmt.Printf("%d\t", root.data)
		root = root.RightTree
	}
}
func isBalanced(root *TreeNode) bool {
	ans := getDepth(root)
	fmt.Println(ans)
	return ans != -1
}
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.LeftTree)
	right := getDepth(root.RightTree)
	return int(math.Max(float64(left), float64(right))) + 1
}
func TreeInit(head *TreeNode) {
	node := new(TreeNode)
	node.data = 2
	head.LeftTree = node
	node = new(TreeNode)
	node.data = 5
	head.RightTree = node
	node = new(TreeNode)
	node.data = 1
	head.LeftTree.LeftTree = node
	node = new(TreeNode)
	node.data = 3
	head.LeftTree.RightTree = node
}
func dfsSearch(root *TreeNode) {
	if root == nil {
		return
	}
	dfsSearch(root.LeftTree)
	fmt.Println(root.data)
	if pre != nil {		//第一次不弄，第一次只把pre设成头
		root.LeftTree = pre		//root的左边 就是前面的节点
		pre.RightTree = root	//前面的节点的后面，就是该节点
	}
	pre = root
	dfsSearch(root.RightTree)
}
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	dfsSearch(root)
	head, tail := root, root
	for head.LeftTree != nil {
		head = head.LeftTree	//left是前面的，这个将head变成刚排好的最小的元素
	}
	for tail.RightTree != nil {
		tail = tail.RightTree	//right是后面的，这个将tail变成刚排好的最大的元素
	}
	head.LeftTree = tail		//这咋还成环了 哈哈哈
	tail.RightTree = head
	return head
}