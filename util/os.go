package util

import (
	"fmt"
	"runtime"
)

func PrintOS() {
	//go switch 不需要break。因为自动添加了，但是如果有fallthrough, 则会向下执行多个label
	//case 无须为常量和整数
	//从上置下，短路
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
