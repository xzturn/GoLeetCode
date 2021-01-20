package intseq

import(
    "math/rand"
    "time"
)

type IntSeqGenerator struct {
}

func NewIntSeqGenerator() *IntSeqGenerator {
    rand.Seed(time.Now().UnixNano())
    return &IntSeqGenerator{}
}

func (g IntSeqGenerator) genKnuth(seq *[]int, n, minv, maxv int) {
    if n <= 0 {
        return
    }
    k := maxv - minv + 1
    if rand.Intn(k) < n {
        (*seq)[n-1] = maxv
        g.genKnuth(seq, n-1, minv, maxv-1)
    } else {
        g.genKnuth(seq, n, minv, maxv-1)
    }
}

func (g IntSeqGenerator) Ordered(n, minv, maxv int) []int {
    seq := make([]int, n)
    g.genKnuth(&seq, n, minv, maxv)
    return seq
}

func (g IntSeqGenerator) Rand(n, minv, maxv int) []int {
    seq := g.Ordered(n, minv, maxv)
    rand.Shuffle(n, func(i, j int) {
        seq[i], seq[j] = seq[j], seq[i]
    })
    return seq
}
