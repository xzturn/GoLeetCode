package GoLeetCode

func zigZagConvert(s string, numRows int) string {
    if (numRows == 1) { return s }
    
    n, step, cur := len(s), (numRows - 1) * 2, 0
    buf := make([]byte, n)
    
    outter := func(i int) {
        for i < n {
            buf[cur] = s[i]
            i += step
            cur++
        }
    }
    
    inner := func(i int) {
        delta := (numRows - i - 1) * 2
        for i < n {
            buf[cur] = s[i]
            cur++
            if k := i + delta; k < n {
                buf[cur] = s[k]
                cur++
            }
            i += step
        }
    }
    
    outter(0)
    for i := 1; i < numRows - 1; i++ {
        inner(i)
    }
    outter(numRows - 1)
    return string(buf)
}
