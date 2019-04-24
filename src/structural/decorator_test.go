package structural

import "testing"

func TestDecorator(t *testing.T) {
	Double := func(n int) int {
		return n * 2
	}

	f := LogDecorate(Double)

	f(5)

}
