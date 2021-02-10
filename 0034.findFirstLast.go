package goleet

func searchRange(nums []int, target int) []int {
    search := func(left, right int) int {
        for left <= right {
            mid := (left + right) >> 1
            switch {
            case nums[mid] < target:
                left = mid + 1
            case nums[mid] > target:
                right = mid - 1
            default:
                return mid
            }
        }
        return -1
    }

    n := len(nums)
    k := search(0, n-1)
    if k == -1 {
        return []int{-1, -1}
    }
    l, r := k, k
    for l-1 >= 0 && nums[l-1] == nums[k] {
        l--
    }
    for r+1 < n && nums[r+1] == nums[k] {
        r++
    }
    return []int{l, r}
}
