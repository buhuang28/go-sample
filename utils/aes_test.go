package utils

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {

	str := "root"

	encrypt, err := AESEncrypt(Str2Bytes(str))
	fmt.Println(encrypt)
	fmt.Println(err)

}
