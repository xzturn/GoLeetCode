package goleet

func removeDuplicates(nums []int) int {
    switch len(nums) {
    case 0:
        return 0
    default:
        i := 0   // current
        for j := 1; j < len(nums); j++ {
            if nums[j] != nums[i] {
                i++
                nums[i] = nums[j]
            }
        }
        return i+1
    }
}

