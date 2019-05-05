package main

import "fmt"

func main() {
	//node1 := &ListNode{3, nil}
	//node2 := &ListNode{2, nil}
	//node3 := &ListNode{0, nil}
	//node4 := &ListNode{-4, nil}
	//head := node1
	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = nodes4
	//node4.Next = node2
	//
	//node := detectCycle(head)
	//fmt.Println(node.Val)
	fmt.Println(isValid("[][][]{}{}{}()()(())"))
}

func isValid(s string) bool {
	// 通过slice实现栈操作
	stack := make([]rune, 0)
	m := map[rune]rune{ //左括号是val，右括号是key
		')': '(',
		']': '[',
		'}': '{', //分段显示，此处必须有逗号
	}

	for _, v := range s {
		if _, ok := m[v]; !ok { // 注意是;不是,
			// 左边括号
			stack = append(stack, v)
		} else if len(stack) != 0 && stack[len(stack)-1] == m[v] { //此处是m[v]，不是v，栈中只存储左括号
			// 右括号
			stack = stack[:len(stack)-1] //出栈
		} else {
			return false
		}
	}
	return len(stack) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输出环的起始位置
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow == fast { // 在此处追上，不代表此处就是环的起始位置
		// slow和fast必须从头开始走，fast不能提前走,
		// A从相遇的位置开始走，每次一步
		// B从head开始，每次一步，相遇时即为环的第一个节点
		slow = head
		for slow != fast {
			slow, fast = slow.Next, fast.Next
		}
		return slow
	}
	return nil
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n&1 == 0 {
		n >>= 1
	}
	if n == 1 {
		return true
	}
	return false
}

func hammingWeight(num uint32) int {
	count := 0
	for num != uint32(0) {
		fmt.Printf("%b\n", num)
		num = num & (num - 1)
		count++
	}
	return count
}

func countBits(num int) []int {
	rnt := []int{}
	for i := int(0); i <= num; i++ {
		rnt = append(rnt, hmweight(i))
	}
	return rnt
}

func hmweight(num int) int {
	var cnt int
	for ; num != 0; num &= (num - 1) {
		cnt++
	}
	return cnt
}
