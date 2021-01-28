package goleet

func myAtoi(s string) int {
    n, i := len(s), 0

    for i < n {
        if s[i] == ' ' {
            i++
        } else {
            break
        }
    }

    flag, res := false, 0
    if i < n {
        if s[i] == '-' {
            flag = true
            i++
        } else if s[i] == '+' {
            i++
        }
    }

    maxv := 1 << 31
    check := func(x int) (int, bool) {
        if x > maxv && flag {
            return -2147483648, false
        } else if x >= maxv && !flag {
            return 2147483647, false
        }
        return x, true
    }

    for i < n {
        if v := int(s[i] - '0'); v < 0 || v > 9 {
            break
        } else {
            res = res * 10 + v
            if x, ok := check(res); !ok {
                return x
            }
        }
        i++
    }

    if flag {
        res = 0 - res
    }

    return res
}
