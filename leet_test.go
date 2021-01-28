package goleet

import (
    "fmt"
    "math/rand"
    "testing"
    "time"

    "github.com/xzturn/xgen"
)

const (
    seqLen = 32
    minVal = 1
    maxVal = 10000
)

var g *xgen.XGenerator = xgen.NewXGenerator()

func prepare0001(n, minv, maxv int) ([]int, int, int, int) {
    seq := g.GenShuffle(seqLen, minVal, maxVal)
    i, j := rand.Intn(seqLen), rand.Intn(seqLen)
    for {
        if i != j {
            break
        }
        j = rand.Intn(seqLen)
    }
    target := seq[i] + seq[j]
    return seq, target, i, j
}

func Test0001(t *testing.T) {
    seq, target, i, j := prepare0001(seqLen, minVal, maxVal)
    fmt.Printf("%v\n%d = %d + %d\ntarget = nums[%d] + nums[%d]\n", seq, target, seq[i], seq[j], i, j)

    ts := time.Now()
    res := twoSum(seq, target)
    te := time.Now()

    fmt.Printf("[%v] %d, %d ==> %v\n", te.Sub(ts), i, j, res)
}

func Benchmark0001(b *testing.B) {
    seq, target, _, _ := prepare0001(seqLen, minVal, maxVal)
    for i := 0; i < b.N; i++ {
        twoSum(seq, target)
    }
}

func Test0006(t *testing.T) {
    s := "PAYPALISHIRING"
    t1 := convert(s, 3)
    t2 := convert(s, 4)
    fmt.Printf("%s [3] --> %s\n", s, t1)
    fmt.Printf("%s [4] --> %s\n", s, t2)
}
