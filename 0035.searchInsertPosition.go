package goleet

func searchInsert(nums []int, target int) int {
    i, j := 0, len(nums)
    for i < j {
        h := int(uint(i+j) >> 1)
        if nums[h] < target {
            i = h + 1
        } else {
            j = h
        }
    }
    return i
}
