func trap(height []int) int {
    var areas, max_l, max_r, l int = 0, 0, 0, 0
    var r int = len(height) - 1
    
    for l < r {
        if height[l] < height[r] {
            if height[l] > max_l {
                max_l = height[l]
            } else {
                areas += max_l - height[l]
            }
            l += 1
        } else {
            if height[r] > max_r {
                max_r = height[r]
            } else {
                areas += max_r - height[r]
            }
            r -= 1
        }
    }
    return areas
}
