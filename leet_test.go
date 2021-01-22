package GoLeetCode

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

func Test0001(t *testing.T) {
    seq := g.GenShuffle(seqLen, minVal, maxVal)

    i, j := rand.Intn(seqLen), rand.Intn(seqLen)
    for {
        if i != j {
            break
        }
        j = rand.Intn(seqLen)
    }
    target := seq[i] + seq[j]
    fmt.Printf("%v\n%d = %d + %d\ntarget = nums[%d] + nums[%d]\n", seq, target, seq[i], seq[j], i, j)

    ts := time.Now()
    res := twoSum(seq, target)
    te := time.Now()

    fmt.Printf("[%v] %d, %d ==> %v\n", te.Sub(ts), i, j, res)
}

func Benchmark0001(b *testing.B) {
}
