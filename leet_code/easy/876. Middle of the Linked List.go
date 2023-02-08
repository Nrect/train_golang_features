package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/middle-of-the-linked-list/description/

//TEST DATA

//func main() {
//	example := make([]easy.ListNode, 10)
//
//	for i := 0; i < 10; i++ {
//		example[i] = easy.ListNode{Val: i}
//
//		if i > 0 {
//			example[i-1].Next = &example[i]
//		}
//	}
//
//	result := make([]easy.ListNode, 0)
//	for _, head := range example {
//		result = append(result, *easy.MiddleNode(&head))
//	}
//
//	fmt.Println(result)
//}

// MiddleNode MySolution
func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil {

		fast = fast.Next

		if fast != nil {
			fast = fast.Next

		} else {
			// fast has reached the end of linked list
			// slow is on the middle point now
			break
		}

		slow = slow.Next
	}

	return slow

}

// other question

func middleNode(head *ListNode) *ListNode {
	listSize := listSize(head)
	tmp := head
	i := 0
	for listSize/2 != i {
		tmp = tmp.Next
		i++
	}
	return tmp
}

func listSize(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + listSize(head.Next)
}
