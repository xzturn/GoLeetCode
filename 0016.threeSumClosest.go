package goleet

import "sort"

func threeSumClosest(nums []int, target int) int {
    abs := func(x int) int {
        if x < 0 {
            return -x
        }
        return x
    }

    sort.Ints(nums)

    delta, res, n := 1<<31 - 1, 0, len(nums)
    for i := range nums {
        l, r, t := i+1, n-1, 0
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            switch {
            case sum > target:
                r--
            case sum < target:
                l++
            default:
                return sum
            }
            if t := abs(sum - target); t < delta {
                res, delta = sum, t
            }
        }
    }

    return res
}
