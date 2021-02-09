package goleet

func search(nums []int, target int) int {
    circled := func(low, high int) int {
        for low <= high {
            mid := (low + high) >> 1
            switch {
            case nums[mid] == target:
                return mid
            case nums[low] <= nums[mid]:  // [low, mid) incr
                if nums[low] <= target && target < nums[mid] {
                    high = mid - 1
                } else {
                    low = mid + 1
                }
            default:  // (mid, high] incr
                if nums[mid] < target && target <= nums[high] {
                    low = mid + 1
                } else {
                    high = mid - 1
                }
            }
        }
        return -1
    }

    return circled(0, len(nums)-1)
}
