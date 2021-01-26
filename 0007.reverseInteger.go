package GoLeetCode

func reverseInteger(x int) int {
    sig, v := 1, x
    if x < 0 {
        sig, v = -1, 0-x
    }

    r := 0
    for v > 0 {
        r = r * 10 + v % 10
        v = v / 10
    }
    
    if r >= 1 << 31 { return 0 }
    
    return r * sig
}
