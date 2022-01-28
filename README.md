# StandardIOSumGo

```golang
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndWriteSum() {
	reader := bufio.NewReader(os.Stdin)
	turns := ReadInt(reader)
	i := 0
	for ; i < turns; i++ {
		nums := ReadInts(reader)
		fmt.Println(nums[0] + nums[1])
	}
}

func ReadLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	result := strings.TrimSpace(line)
	return result
}

func ReadInt(reader *bufio.Reader) int {
	line := ReadLine(reader)
	return ParseInt(line)
}

func ReadInts(reader *bufio.Reader) []int {
	line := ReadLine(reader)
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
```