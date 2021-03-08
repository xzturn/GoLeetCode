package goleet

func plusOne(digits []int) []int {
    n := len(digits)
    res, carry := make([]int, n), 0
    if res[n-1] = digits[n-1] + 1; res[n-1] >= 10 {
        res[n-1] -= 10
        carry = 1
    }
    for i := n-2; i >=0; i-- {
        res[i] = digits[i] + carry
        if res[i] >= 10 {
            res[i] -= 10
            carry = 1
        } else {
            carry = 0
        }
    }
    if carry == 1 {
        res = append([]int{1}, res...)
    }
    return res
}
