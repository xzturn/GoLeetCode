package goleet

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 1 <= n <= sz
    p, i, q := head, 1, head
    for ; i <= n; i++ {
        p = p.Next
    }
    if p == nil {
        head = head.Next
        return head
    }
    for p.Next != nil {
        p = p.Next
        q = q.Next
    }
    if p == head {
        return nil
    }
    q.Next = q.Next.Next
    return head
}
