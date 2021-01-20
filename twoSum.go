package main

import (
    "flag"
    "fmt"
    "math/rand"
    "os"
    "runtime"
    "time"

    "xzturn/intseq"
)

func twoSum(nums []int, target int) []int {
    res := make([]int, 2)
    aid := make(map[int] int)
    for i, x := range nums {
        if j, ok := aid[target - x]; ok {
            res[0], res[1] = i, j
            return res
        }
        aid[x] = i
    }
    return nil
}

var seqLen *int = flag.Int("n", 32, "the rand int sequence length")
var minVal *int = flag.Int("min", 1, "the min value of the sequence")
var maxVal *int = flag.Int("max", 10000, "the max value of the sequence")

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()

    if *maxVal - *minVal + 1 < *seqLen || *seqLen < 1 {
        fmt.Fprintf(os.Stderr, "Bad paramerter: n = %d, minv = %d, maxv = %d\n", *seqLen, *minVal, *maxVal)
        os.Exit(1)
    }

    g := intseq.NewIntSeqGenerator()
    seq := g.Rand(*seqLen, *minVal, *maxVal)

    i, j := rand.Intn(*seqLen), rand.Intn(*seqLen)
    for {
        if i != j {
            break
        }
        j = rand.Intn(*seqLen)
    }
    target := seq[i] + seq[j]
    fmt.Printf("%v\n%d = %d + %d\ntarget = nums[%d] + nums[%d]\n", seq, target, seq[i], seq[j], i, j)

    ts := time.Now()
    res := twoSum(seq, target)
    te := time.Now()

    fmt.Printf("[%v] %d, %d ==> %v\n", te.Sub(ts), i, j, res)
}
