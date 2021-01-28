package goleet

// ListNode ...
type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    fp := func(p **ListNode) int {
        x := (*p).Val / 10
        (*p).Val %= 10
        return x
    }

    mv := func(p **ListNode) int {
        if *p != nil {
            if *p = (*p).Next; *p != nil {
                return (*p).Val
            }
        }
        return 0
    }

    l := &ListNode{l1.Val + l2.Val, nil}
    x := fp(&l)
    k, v1, v2 := l, 0, 0
    for {
        v1, v2 = mv(&l1), mv(&l2)
        if l1 == nil && l2 == nil { break }
        k.Next = &ListNode{v1 + v2 + x, nil}
        k = k.Next
        x = fp(&k)
    }
    if x != 0 {
        k.Next = &ListNode{x, nil}
    }
    return l
}
