package main

import (
    "flag"
    "fmt"
    //"math/rand"
    //"os"
    "runtime"
    //"time"
)


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

var minLen *int = flag.Int("min", 10, "the min length of the string")
var maxLen *int = flag.Int("max", 100, "the max length of the string")

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
    flag.Parse()
    x := int('a')
    var c rune = 99
    t := [256]int{}
    fmt.Printf("%d, %c\n%v\n%d\n", x, c, t, t['x'])
}
