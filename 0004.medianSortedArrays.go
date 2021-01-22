package GoLeetCode

func findMedianSortedArrays(nums1 []int, nums2 int[]) float64 {
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

    max := func(x, y int) int { if x > y { return x } else { return y } }
    min := func(x, y int) int { if x < y { return x } else { return y } }

    left, right := 0, 0
    if mid1 == 0 {
        left = nums2[mid2 - 1]
    } else if mid2 == 0 {
        left = nums1[mid1 - 1]
    } else {
        left = max(nums1[mid1 -1], nums2[mid2 - 1])
    }
    if (n1 + n2) & 1 == 1 {
        return float64(left)
    }

    if mid1 == n1 {
        right = nums2[mid2]
    } else if mid2 == n2 {
        right = nums1[mid1]
    } else {
        right = min(nums1[mid1], nums2[mid2])
    }
    return float64(left + right) / 2
}
