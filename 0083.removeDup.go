package goleet

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    p, q := head, head
    for p != nil {
        q = p.Next
        for q != nil && q.Val == p.Val {
            q = q.Next
        }
        p.Next = q
        p = q
    }
    return head
}
