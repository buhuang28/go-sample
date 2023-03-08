package caches

import (
	"fmt"
	"testing"
)

func TestMid(t *testing.T) {
	for i := 0; i < 6588888; i++ {
		fmt.Println(GetMid())
	}
}
