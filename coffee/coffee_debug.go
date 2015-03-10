// +build debug

package coffee

import (
	"fmt"
)

func NewCoffee() Coffee {
	return Coffee{"debug"}
}

func (c Coffee) Println() {
	fmt.Println("[DEBUG]", c.kind)
}
