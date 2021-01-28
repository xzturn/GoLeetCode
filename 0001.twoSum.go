package goleet

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
