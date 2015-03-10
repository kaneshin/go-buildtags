// +build !debug

package coffee

import (
	"fmt"
)

func NewCoffee() Coffee {
	return Coffee{"prod"}
}

func (c Coffee) Println() {
	fmt.Println(c.kind)
}
