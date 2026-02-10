package common

import (
	"fmt"
	"testing"
)

func TestToFloat64Round2(t *testing.T) {
	res, bool := ToFloat64Round2("1.254")
	fmt.Println(res, bool)
}
