package GoLeetCode

func convert(s string, numRows int) string {
    if (numRows == 1) { return s }
    
    outter := func(s string, n, r, i, step int, buf *[]byte, cur int) int {
        for i < n {
            (*buf)[cur] = s[i]
            i += step
            cur++
        }
        return cur
    }
    
    inner := func(s string, n, r, i, step int, buf *[]byte, cur int) int {
        delta := (r - i - 1) * 2
        for i < n {
            (*buf)[cur] = s[i]
            cur++
            if k := i + delta; k < n {
                (*buf)[cur] = s[k]
                cur++
            }
            i += step
        }
        return cur
    }
    
    n, step := len(s), (numRows - 1) * 2
    buf := make([]byte, n)
    cur := outter(s, n, numRows, 0, step, &buf, 0)
    for i := 1; i < numRows - 1; i++ {
        cur = inner(s, n, numRows, i, step, &buf, cur)
    }
    outter(s, n, numRows, numRows - 1, step, &buf, cur)
    return string(buf)
}
