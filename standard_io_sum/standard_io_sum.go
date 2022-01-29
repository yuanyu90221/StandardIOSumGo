package standard_io_sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndWriteSum() string {
	reader := bufio.NewReader(os.Stdin)
	turns := readInt(reader)
	i := 0
	result := ""
	for ; i < turns; i++ {
		nums := readInts(reader)
		result += fmt.Sprintf("%d\n", nums[0]+nums[1])
	}
	return result
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	result := strings.TrimSpace(line)
	return result
}

func readInt(reader *bufio.Reader) int {
	line := readLine(reader)
	return parseInt(line)
}

func readInts(reader *bufio.Reader) []int {
	line := readLine(reader)
	lines := strings.Split(line, " ")
	nums := []int{}
	for _, num := range lines {
		nums = append(nums, parseInt(num))
	}
	return nums
}

func parseInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}
