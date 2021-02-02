package goleet

func removeDuplicates(nums []int) int {
    n := len(nums)
    switch n {
    case 0:
        return 0
    default:
        i := 0   // current
        for j := 1; j < n; j++ {
            if nums[j] != nums[i] {
                i++
                nums[i] = nums[j]
            }
        }
        return i+1
    }
}

