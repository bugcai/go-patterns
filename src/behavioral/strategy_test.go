package behavioral

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	add := Operation{Addition{}}
	result := add.Operate(3, 5) // 8
	fmt.Println("The result of addition is ", result)

	mult := Operation{Multiplication{}}
	result = mult.Operate(3, 5) // 15
	fmt.Println("The result of multiplication is ", result)
}
