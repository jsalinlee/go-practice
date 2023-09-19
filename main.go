package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/**
	 * Definition for singly-linked list.
	 * type ListNode struct {
	 *     Val int
	 *     Next *ListNode
	 * }
	 */

	linkedList := &ListNode{1, nil}
	runner := linkedList
	for i := 2; i < 6; i++ {
		runner.Next = &ListNode{i, nil}
		runner = runner.Next
	}

	printRunner := linkedList
	for printRunner != nil {
		fmt.Println(printRunner)
		printRunner = printRunner.Next
	}
	fmt.Println("----------------------------")
	newList := reverseList(linkedList)
	for newList != nil {
		fmt.Println(newList)
		newList = newList.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	stack := []*ListNode{}
	runner := head
	for runner.Next != nil {
		stack = append(stack, runner)
		runner = runner.Next
	}
	newHead := runner

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		runner.Next = node
		node.Next = nil
		runner = runner.Next
	}

	return newHead
}
