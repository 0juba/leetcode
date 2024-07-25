package container_with_most_water

func maxArea(height []int) int {
	maxAr := 0
	for i, j := 0, len(height)-1; i < j; {
		h1, h2 := height[i], height[j]
		minH := h1
		if h2 < h1 {
			minH = h2
		}

		area := minH * (j - i)
		if maxAr < area {
			maxAr = area
		}

		if h1 > h2 {
			j--
		} else {
			i++
		}
	}

	return maxAr
}
