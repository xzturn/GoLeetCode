package goleet

func strStr(haystack string, needle string) int {
    n, k := len(needle), len(haystack)
    if n == 0 && k >= 0 {
        return 0
    }
    for i := 0; i <= k-n; i++ {
        j := 0
        for j < n {
            if haystack[i+j] != needle[j] {
                break
            }
            j++
        }
        if j == n {
            return i
        }
    }
    return -1
}
