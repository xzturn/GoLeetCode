package goleet

func letterCombinations(digits string) []string {
    return combo(digits, 0)
}

func n2s(c uint8) []string {
    switch c {
    case '2':
        return []string{"a", "b", "c"}
    case '3':
        return []string{"d", "e", "f"}
    case '4':
        return []string{"g", "h", "i"}
    case '5':
        return []string{"j", "k", "l"}
    case '6':
        return []string{"m", "n", "o"}
    case '7':
        return []string{"p", "q", "r", "s"}
    case '8':
        return []string{"t", "u", "v"}
    case '9':
        return []string{"w", "x", "y", "z"}
    }
    return nil
}

func combo(digits string, start int) []string {
    n := len(digits)
    if start > n - 1 {
        return nil
    }
    cur := n2s(digits[start])
    if tmp := combo(digits, start + 1); tmp != nil {
        n0, n1 := len(cur), len(tmp)
        res := make([]string, n0 * n1)
        for i := 0; i < n0; i++ {
            for j := 0; j < n1; j++ {
                res[i*n1+j] = cur[i] + tmp[j]
            }
        }
        return res
    }
    return cur
}
