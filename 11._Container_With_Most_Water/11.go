package containerwithmostwater

// 11. Container With Most Water

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

func maxArea(height []int) int {
	// use 2 pointers to solve this problem
	// if [1 2 3 4]
	// left is 1
	// right is 4
	l, r := 0, len(height)-1
	res := 0

	// if left is not yet intersect the right continue the iteration
	for l < r {
		//calculate the max area
		hl, hr := height[l], height[r]
		w := r - l
		h := min(hl, hr)
		area := w * h

		// compare the area
		res = max(area, res)

		// if left height is lower than right height
		// increase the l
		if hl < hr {
			l++
		} else {
			// otherwise decrease th r
			r--
		}
	}

	return res
}
