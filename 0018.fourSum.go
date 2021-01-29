package goleet

import "sort"

func fourSum(nums []int, target int) [][]int {
    n, res := len(nums), [][]int{}
    if n < 4 {
        return res
    }
    
    sort.Ints(nums)
   
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
    
    select2 := func(i, j, l, r int) {
        t := nums[i] + nums[j]
        for l < r {
            v := t + nums[l] + nums[r]
            switch {
            case v < target:
                l++
            case v > target:
                r--
            default:
                res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
                l, r = next(l, r)
            }
        }
    }

    next2 := func(i, j int) (int, int) {
        j++
        for j < n && nums[j] == nums[j-1] {
            j++
        }
        if j < n {
            return i, j
        }
        
        i++
        for i < n && nums[i] == nums[i-1] {
            i++
        }
        return i, i+1
    }
    
    i, j := 0, 1
    for i < n && j < n {
        select2(i, j, j+1, n-1)
        i, j = next2(i, j)
    }
    
    return res
}
