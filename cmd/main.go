package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/malinatrash/yadro-go-intern/task"
)

func readContainers(n int, reader *bufio.Scanner) [][]int {
	containers := make([][]int, n)
	for i := 0; i < n && reader.Scan(); i++ {
		line := strings.Fields(reader.Text())
		containers[i] = make([]int, len(line))
		for j, valStr := range line {
			val, err := strconv.Atoi(valStr)
			if err != nil {
				panic(err)
			}
			containers[i][j] = val
		}
	}
	return containers
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	containers := readContainers(n, scanner)

	result := task.CanSortBalls(containers)
	fmt.Println(result)
}
