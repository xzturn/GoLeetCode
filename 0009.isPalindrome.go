package GoLeetCode

func isPalindrome(x int) bool {
    if x < 0 { return false }

    y, z := x, 0
    for ; x > 0; x /= 10 {
        z = z * 10 + (x % 10)
    }
    return y == z
}
