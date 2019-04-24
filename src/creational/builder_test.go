package creational

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := NewCarBuilder()
	car := builder.Color(RedColor).TopSpeed(50).Wheels(SportsWheels).Build()

	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}
