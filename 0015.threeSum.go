package goleet

import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res, n := make([][]int, 0), len(nums)

    next := func(l, r int) (int, int) {
        for l < r {
            switch {
            case nums[l] == nums[l+1]:
                l++
            case nums[r] == nums[r-1]:
                r--
            default:
                l++
                r--
                return l, r
            }
        }
        return l, r
    }

    for i := range nums {
        if nums[i] > 0 {
            break
        }

        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l, r := i+1, n-1
        for l < r {
            v := nums[i] + nums[l] + nums[r]
            switch {
            case v < 0:
                l++
            case v > 0:
                r--
            default:
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l, r = next(l, r)
            }
        }
    }

    return res
}
