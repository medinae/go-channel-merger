package merger

import (
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	x := chanOf(1, 2, 3, 4, 5)
	y := chanOf(34, 345)

	m := NewMerger()
	merged := m.Merge(x, y)

	c := 0
	expected := 7

	for range merged {
		c++
	}

	if c != expected {
		t.Errorf("Merged chan values count was incorrect, got: %d, want: %d.", c, expected)
	}
}

func chanOf(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
			waitTwoSec()
		}
	}()

	return ch
}

func waitTwoSec() {
	time.Sleep(
		time.Duration(1000) * time.Millisecond,
	)
}
