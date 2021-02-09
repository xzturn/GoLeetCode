package goleet

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    p, q := head, head.Next
    for q != nil {
        p.Val, q.Val = q.Val, p.Val
        p = q.Next
        if p != nil {
            q = p.Next
        } else {
            q = nil
        }
    }

    return head
}
