package task

import (
	"sort"
)

func CanSortBalls(containers [][]int) string {
	colorCounts := make([]int, len(containers))
	containerCounts := make([]int, len(containers))
	for i, container := range containers {
		for _, count := range container {
			containerCounts[i] += count
		}
		for index, count := range container {
			colorCounts[index] += count
		}
	}
	sort.Ints(containerCounts)
	sort.Ints(colorCounts)

	if isEqual(colorCounts, containerCounts) {
		return "yes"
	} else {
		return "no"
	}
}

func isEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
