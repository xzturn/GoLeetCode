package goleet

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1, n2 := len(nums1), len(nums2)
    if n1 > n2 {
        return findMedianSortedArrays(nums2, nums1)
    }

    low, high, k, mid1, mid2 := 0, n1, (n1 + n2 + 1) >> 1, 0, 0
    for low <= high {
        // nums1: [0, mid1) | [mid1, n1)
        // nums2: [0, mid2) | [mid2, n2)
        mid1 = low + (high - low) >> 1
        mid2 = k - mid1
        if mid1 > 0 && nums1[mid1 - 1] > nums2[mid2] {
            high = mid1 - 1
        } else if mid1 != n1 && nums1[mid1] < nums2[mid2 - 1] {
            low = mid1 + 1
        } else {
            break
        }
    }

    left, right := 0, 0
    if mid1 == 0 {
        left = nums2[mid2 - 1]
    } else if mid2 == 0 {
        left = nums1[mid1 - 1]
    } else if nums1[mid1 - 1] > nums2[mid2 - 1] {
        left = nums1[mid1 -1]
    } else {
        left = nums2[mid2 - 1]
    }
    if (n1 + n2) & 1 == 1 {
        return float64(left)
    }

    if mid1 == n1 {
        right = nums2[mid2]
    } else if mid2 == n2 {
        right = nums1[mid1]
    } else if nums1[mid1] < nums2[mid2] {
        right = nums1[mid1]
    } else {
        right = nums2[mid2]
    }
    return float64(left + right) / 2
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
    i1, i2, n1, n2 := 0, 0, len(nums1), len(nums2)
    i, n := 0, n1 + n2
    nums := make([]int, n)
    for i1 < n1 && i2 < n2 {
        if nums1[i1] < nums2[i2] {
            nums[i] = nums1[i1]
            i1++
        } else {
            nums[i] = nums2[i2]
            i2++
        }
        i++
    }
    for i1 < n1 {
        nums[i] = nums1[i1]
        i++
        i1++
    }
    for i2 < n2 {
        nums[i] = nums2[i2]
        i++
        i2++
    }
    k := n / 2
    if n % 2 == 1 {
        return float64(nums[k])
    }
    return float64(nums[k] + nums[k-1]) / 2
}
