package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getNumber(temp *ListNode) *big.Int {
	var result string

	for temp != nil {
		result = strconv.Itoa(temp.Val) + result
		temp = temp.Next
	}

	number := new(big.Int)
	number.SetString(result, 10)
	return number
}

func changeToListNode(number *big.Int) *ListNode {
	numberStr := number.String()
	arr := []rune(numberStr)

	var result, temp *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		val := int(arr[i] - '0')
		if result == nil {
			result = &ListNode{Val: val}
			temp = result
		} else {
			temp.Next = &ListNode{Val: val}
			temp = temp.Next
		}
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Num := getNumber(l1)
	l2Num := getNumber(l2)

	sum := new(big.Int).Add(l1Num, l2Num)

	return changeToListNode(sum)
}

func arrayToListNode(arr []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range arr {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

func main() {
	l1 := arrayToListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := arrayToListNode([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)
	printList(result)
}
