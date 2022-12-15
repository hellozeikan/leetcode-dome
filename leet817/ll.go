package leet817

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	res := 0
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] = 1
	}
	for head != nil {
		if _, ok := m[head.Val]; ok {
			for head != nil {
				if _, ok := m[head.Val]; ok {
					head = head.Next
				} else {
					break
				}
			}
			res++
		} else {
			head = head.Next
		}
	}
	return res
}

// 输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
// 输出：5
// 解释：可以采摘 [1,2,1,1,2] 这五棵树。
func totalFruit(fruits []int) int {
	s := []int{}
	for _, elem := range fruits {
		s = append(s, elem)
	}
	len := len(fruits)
	if len == 1 {
		return 1
	}
	res := 0
	i, j := 0, 0
	for i < len && i <= j && j < len {
		if isTwoValue(s[i:j]) {
			res = max(res, j-i+1)
			j++
		} else {
			i++
		}
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func isTwoValue(arr []int) bool {
	a, b := -1, -1
	for _, v := range arr {
		if a != b && a != v && b != v {
			return false
		}
		if a != v {
			b = v
		}
		if a == -1 {
			a = v
		}

	}
	return true
}

func countStudents(students []int, sandwiches []int) int {
	count := [2]int{}
	for _, val := range students {
		count[val]++
	}
	for _, val := range sandwiches {
		if count[val] == 0 && val == 1 {
			return count[0]
			// 实际上的写法就是拿另一个索引的值
		} else if count[val] == 0 {
			return count[1]
		}
		count[val]--
	}
	return 0
}
