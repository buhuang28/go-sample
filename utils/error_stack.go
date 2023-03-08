package utils

import (
	"bytes"
	"fmt"
	"runtime"
)

func PrintStackTrace(err interface{}) string {
	buffer := new(bytes.Buffer)
	fmt.Fprintf(buffer, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buffer, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buffer.String()
}
