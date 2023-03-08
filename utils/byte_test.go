package utils

import (
	"fmt"
	"testing"
)

func TestBytes2Int(t *testing.T) {
	b := []byte{1, 2}
	fmt.Println(LitBytes2Int(b))
}

func TestNilMap(t *testing.T) {
	a := make(map[string]string)
	var b map[string]string
	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
