package goleet

func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }

    var freq [256]int
    res, cur, left, right := 0, 0, 0, 0
    for left < n {
        if right < n && freq[s[right]] == 0 {
            freq[s[right]]++
            right++
        } else {
            freq[s[left]]--
            left++
        }
        if cur = right - left; res < cur {
            res = cur
        }
    }
    return res
}
