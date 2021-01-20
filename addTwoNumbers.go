package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    fp := func(p **ListNode) int {
        x := 0
        if (*p).Val > 9 {
            (*p).Val -= 10
            x = 1
        }
        return x
    }

    l := &ListNode{l1.Val + l2.Val, nil}
    x := fp(&l)
    k := l
    for {
        l1, l2 = l1.Next, l2.Next
        if l1 == nil { break }
        if l2 == nil { break }
        k.Next = &ListNode{l1.Val + l2.Val + x, nil}
        k = k.Next
        x = fp(&k)
    }
    for l1 != nil {
        k.Next = &ListNode{l1.Val + x, nil}
        k = k.Next
        x = fp(&k)
        l1 = l1.Next
    }
    for l2 != nil {
        k.Next = &ListNode{l2.Val + x, nil}
        k = k.Next
        x = fp(&k)
        l2 = l2.Next
    }
    if x != 0 {
        k.Next = &ListNode{x, nil}
    }
    return l
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
}
