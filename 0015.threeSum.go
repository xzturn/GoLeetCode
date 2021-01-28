package GoLeetCode

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

func threeSum2(nums []int) [][]int {
    res := make([][]int, 0)

    aid := make(map[int] int)
    for _, n := range nums {
        if _, ok := aid[n]; !ok {
            aid[n] = 0
        }
        aid[n]++
    }

    ns := make([]int, 0)
    for k, v := range aid {
        if v >= 2 {
            if k == 0 {
                if v >= 3 {
                    res = append(res, []int{0, 0, 0})
                }
                ns = append(ns, 0)
            } else {
                ns = append(ns, k, k)
            }
        } else {
            ns = append(ns, k)
        }
    }

    sort.Ints(ns)
    n := len(ns)

    bsearch := func(x, left, right int) int {
        low, high, mid := left-1, right+1, 0
        for low+1 < high {
            mid = (low + high) / 2
            if ns[mid] < x {
                low = mid
            } else {
                high = mid
            }
        }
        if high > right || ns[high] != x {
            return -1
        }
        return high
    }

    findrm := func(l, r int) {
        for r > l + 1 {
            if r < n - 1 && ns[r] == ns[r+1] {
                r--
            } else {
                x := 0 - ns[l] - ns[r]
                if x < 0 && ns[l] >=0 { return }
                if x > 0 && ns[r] <=0 { return }
                if m := bsearch(x, l+1, r-1); m != -1 {
                    res = append(res, []int{ns[l], x, ns[r]})
                }
                r--
            }
        }
    }

    findrm(0, n-1)
    for l := 1; l < n-2; {
        if ns[l] > ns[l-1] {
            findrm(l, n-1)
        }
        l++
    }
    return res
}
