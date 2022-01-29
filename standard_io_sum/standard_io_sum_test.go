package standard_io_sum

import (
	"testing"

	"github.com/rhysd/go-fakeio"
)

func TestReadAndWriteSum(t *testing.T) {
	fake := fakeio.Stdin("4\n1 5\n314 15\n-99 99\n123 987\n").Stderr().Stdout()
	want := "6\n329\n0\n1110\n"
	got := ReadAndWriteSum()
	_, err := fake.String()
	fake.Restore()
	if err != nil {
		t.Errorf("%v", err)
	}
	if got != want {
		t.Errorf("ReadAndWriteSum() = %v, want %v", got, want)
	}
}
