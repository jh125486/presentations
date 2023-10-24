package main

import (
	"fmt"
	"runtime"
)

func main() {
	// START OMIT
	switch os := runtime.GOOS; os { // HL
	case "darwin":
		fmt.Println("macOS", runtime.Version())
	case "linux":
		fmt.Println("GNU/Linux", runtime.Version())
	case "windows":
		fmt.Println("Microsoft Windows", runtime.Version())
	default: // freebsd, openbsd, plan9...
		fmt.Printf("%s. %v", os, runtime.Version())
	} // HL
	// END OMIT
}
