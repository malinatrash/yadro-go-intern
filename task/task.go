package task

func CanSortBalls(containers [][]int) string {
	n := len(containers)

	colorCounts := make(map[int]int)
	for _, container := range containers {
		for _, color := range container {
			colorCounts[color]++
		}
	}

	for _, count := range colorCounts {
		if count%n != 0 {
			return "no"
		}
	}

	for i := 0; i < n; i++ {
		if len(containers[i]) != n {
			return "no"
		}
	}

	for _, container := range containers {
		seenColors := make(map[int]bool)
		for i, _ := range container {
			seenColors[i] = true
		}
		if len(seenColors) != n {
			return "no"
		}
	}

	return "yes"
}
