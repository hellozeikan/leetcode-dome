package main

import (
	"fmt"
)

func main() {
	head := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{}}}}}
	temp := reverseKGroup(&head, 2, &ListNode{})
	fmt.Println(temp.Val)
	fmt.Println(temp.Next.Val)
	fmt.Println(temp.Next.Next.Val)
	fmt.Println(temp.Next.Next.Next.Val)
	// arrayStringsAreEqual([]string{"qwe"}, []string{"qwer"})
}
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	arr1 := []byte{}
	arr2 := []byte{}
	for _, val := range word1 {
		arr1 = append(arr1, []byte(val)...)
	}
	for _, val := range word2 {
		arr2 = append(arr2, []byte(val)...)
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int, succ *ListNode) *ListNode {
	if k == 1 {
		succ = head.Next
		return head
	}
	last := reverseKGroup(head.Next, k-1, succ)
	head.Next.Next = head
	head = succ
	return last
}
