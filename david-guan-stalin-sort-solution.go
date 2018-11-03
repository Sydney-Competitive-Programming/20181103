package main

import (
	"fmt"
	"math"
	"sort"
)

type node struct {
	left, right *node
	val         int
}

func (root *node) insert(newOne *node) {
	if root == nil {
		return
	}
	// Insert right
	if newOne.val >= root.val {
		if root.right == nil {
			root.right = newOne
			return
		}
		root.right.insert(newOne)
		return
	}
	// Insert left
	if root.left == nil {
		root.left = newOne
		return
	}
	root.left.insert(newOne)
}

func (root *node) getRightChainLength() int {
	res := 0
	for root != nil {
		res, root = res+1, root.right
	}
	return res
}

func (root *node) getScore() int {
	if root == nil {
		return 0
	}
	rightChainLen := float64(root.getRightChainLength())
	choiceTwo := float64(root.left.getScore())
	choiceThree := float64(root.right.getScore() + 1)
	return int(math.Max(math.Max(rightChainLen, choiceTwo), choiceThree))
}

func (root *node) getScoreList() []int {
	res, queue := []int{}, []*node{root}
	if root == nil {
		return res
	}
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		if task.left != nil {
			queue = append(queue, task.left)
		}
		if task.right != nil {
			queue = append(queue, task.right)
		}
		res = append(res, task.getScore())
	}
	return res
}

// Just for test the tree is correctly build
// func (root *node) displayTree() {
// 	if root == nil {
// 		return
// 	}
// 	root.left.displayTree()
// 	root.right.displayTree()
// 	fmt.Println(root.val)
// }

func solve(src []int) []int {
	if len(src) == 0 {
		return []int{}
	}
	// Build the tree
	root := node{val: src[0]}
	for i := 1; i < len(src); i++ {
		newNode := node{val: src[i]}
		root.insert(&newNode)
	}
	// Calc scores:
	scoreList := root.getScoreList()
	// Find the value
	sort.Ints(scoreList)
	if len(scoreList) == 1 {
		return scoreList
	}
	return scoreList[len(scoreList)-2:]
}

func main() {
	cases := [][]int{
		[]int{1, 2, 4, 5, 3, 6, 6},
		[]int{19, 2},
		[]int{3, 3, 4, 3},
		[]int{10},
		[]int{1, 2, 4, 9},
		[]int{1, 90, 2, 3, 4, 5},
		[]int{1, 90, 91, 2, 3, 4, 5},
		[]int{3, 4, 5, 3, 1, 2, 7, 8, 9, 10, 11},
		[]int{1, 21, 41, 61, 12, 6, 88, 13, 14, 15, 16, 17, 18, 7, 8, 9, 10},
		[]int{2, 1, 4, 3, 5, 6, 8, 7},
	}
	for _, c := range cases {
		fmt.Println("Case:", c)
		list := solve(c)
		if len(list) == 1 {
			fmt.Println("Only one sorted arr, the length is:", list[0])
		} else {
			fmt.Println("First:", list[1], "Second:", list[0])
		}
	}
}
