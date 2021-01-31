package goleet

func longestCommonPrefix(strs []string) string {
    n := len(strs)

    walk := func() string {
        k := len(strs[0])
        buf, cur := make([]byte, k), 0
        for cur < k {
            t := strs[0][cur]
            for i := 1; i < n; i++ {
                if cur >= len(strs[i]) || strs[i][cur] != t {
                    return string(buf[0:cur])
                }
            }
            buf[cur] = t
            cur++
        }
        return string(buf[0:cur])
    }

    switch n {
    case 0:
        return ""
    case 1:
        return strs[0]
    default:
        return walk()
    }
}
