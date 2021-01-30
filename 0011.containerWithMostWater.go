package goleet

func maxArea(height []int) int {
    n, ma := len(height), 0
    l, r, h := 0, n - 1, 0
    for l < r {
        w := r - l
        switch {
        case height[l] < height[r]:
            h = height[l]
            l++
        case height[l] > height[r]:
            h = height[r]
            r--
        default:
            h = height[l]
            l++
            r--
        }
        if t := w * h; t > ma {
            ma = t
        }
    }

    return ma
}
