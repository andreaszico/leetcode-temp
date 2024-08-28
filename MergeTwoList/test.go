package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Creating first linked list: 1 -> 2 -> 3 -> nil
	nodeA1 := &ListNode{Val: 1}
	nodeA2 := &ListNode{Val: 2}
	nodeA3 := &ListNode{Val: 3}
	nodeA1.Next = nodeA2
	nodeA2.Next = nodeA3

	// Creating second linked list: 4 -> 5 -> 6 -> nil
	nodeB1 := &ListNode{Val: 1}
	nodeB2 := &ListNode{Val: 3}
	nodeB3 := &ListNode{Val: 5}
	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3

	// Printing the linked lists
	fmt.Println("First Linked List:")
	printLinkedList(nodeA1)

	fmt.Println("Second Linked List:")
	printLinkedList(nodeB1)
	println("=============================")

	result := mergeSorted(nodeA1, nodeB1)
	fmt.Println("Result Linked List:")
	printLinkedList(result)
}

func mergeSorted(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	result := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			result.Next = list1
			list1 = list1.Next
		} else {
			result.Next = list2
			list2 = list2.Next
		}
		result = result.Next
	}

	if list1 != nil {
		result.Next = list1
	} else if list2 != nil {
		result.Next = list2
	}

	return dummy.Next
}

// Helper function to print the linked list
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
