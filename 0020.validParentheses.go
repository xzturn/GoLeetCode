package goleet

func isValid(s string) bool {
    // use stack
    top, n := 0, len(s)
    st := make([]uint8, n)
    push := func(c uint8) {
        st[top] = c
        top++
    }
    pop := func() (uint8, bool) {
        top--
        if top >= 0 {
            return st[top], true
        }
        return 0, false
    }
    check := func(c uint8) bool {
        if t, ok := pop(); ok && t == c {
            return true
        }
        return false
    }

    for i := 0; i < n; i++ {
        switch s[i] {
        case ')':
            if !check('(') { return false }
        case ']':
            if !check('[') { return false }
        case '}':
            if !check('{') { return false }
        default:  // '(', '[', '{'
            push(s[i])
        }
    }

    return top <= 0
}
