package standard_io_sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndWriteSum() {
	reader := bufio.NewReader(os.Stdin)
	turns := readInt(reader)
	i := 0
	for ; i < turns; i++ {
		nums := readInts(reader)
		fmt.Println(nums[0] + nums[1])
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	result := strings.TrimSpace(line)
	return result
}

func readInt(reader *bufio.Reader) int {
	line := readLine(reader)
	return ParseInt(line)
}

func readInts(reader *bufio.Reader) []int {
	line := readLine(reader)
	lines := strings.Split(line, " ")
	nums := []int{}
	for _, num := range lines {
		nums = append(nums, ParseInt(num))
	}
	return nums
}

func ParseInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}
