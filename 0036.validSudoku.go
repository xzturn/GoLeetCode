package goleet

// [15, 10] reserved, [9, 1] indicate digits
// [0] indicate if the value is a given value
type bitmap9 uint16

func (b *bitmap9) set(i int) {
    *b |= 1 << i
}

func (b *bitmap9) clr(i int) {
    *b &= ^(1 << i)
}

func (b bitmap9) chk(i int) bool {
    return b & (1 << i) != 0
}

func (b bitmap9) empty() bool {
    return b & 0x3fe == 0
}

func isValidSudoku(board [][]byte) bool {
    getv := func(i, j int) int {
        // ['1', '9'] -> [1, 9]
        if board[i][j] >= '1' && board[i][j] <= '9' {
            return int(board[i][j] - '0')
        }
        return 0
    }
    
    subidx := func(i, j int) int {
        return (i/3) * 3 + (j/3)
    }

    status := make([][]bitmap9, 9)
    for i := 0; i < 9; i++ {
        status[i] = make([]bitmap9, 9)
    }    
    row, col, sub := make([]bitmap9, 9), make([]bitmap9, 9), make([]bitmap9, 9)

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if v := getv(i, j); v != 0 {
                // process given numbers
                if row[i].chk(v) {
                    // duplicate v in a row
                    return false
                }
                if col[j].chk(v) {
                    // duplicate v in a col
                    return false
                }
                k := subidx(i, j)
                if sub[k].chk(v) {
                    // duplicate v in a subbox
                    return false
                }
                row[i].set(v)         // v in row
                col[j].set(v)         // v in col
                sub[k].set(v)         // v in subbox
                status[i][j].set(v)   // v is pos(i,j)
                status[i][j].set(0)   // given flag
            }
        }
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if !status[i][j].chk(0) {
                // process blank positions
                status[i][j] = ^row[i] & ^col[j] & ^sub[subidx(i,j)]
                if status[i][j].empty() {
                    return false
                }
            }
        }
    }
 
    return true
}
