package GoLeetCode

func longestPalindrome(s string) string {
    palindrome := func(s string, n, l, r int) (int, int, int) {
        for l >= 0 && r < n  && s[l] == s[r] {
            l--
            r++
        }
        return l+1, r, r-l-1
    }
    
    n, l, r, k := len(s), 0, 0, 0
    for i := 0; i < n; i++ {
        l1, r1, k1 := palindrome(s, n, i, i)
        l2, r2, k2 := palindrome(s, n, i, i+1)
        if k < k1 { l, r, k = l1, r1, k1 }
        if k < k2 { l, r, k = l2, r2, k2 }
    }
    return s[l:r]
}
