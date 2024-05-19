package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/malinatrash/yadro-go-intern/task"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]int, n)
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		for j := 0; j < n; j++ {
			containers[i][j], _ = strconv.Atoi(parts[j])
		}
	}

	result := task.CanSortBalls(containers)
	fmt.Println(result)
}
