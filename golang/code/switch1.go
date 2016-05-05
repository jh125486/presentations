package main

import (
	"fmt"
	"runtime"
)

func main() {
	// START OMIT
	switch os := runtime.GOOS; os { // HL
	case "darwin":
		fmt.Println("OS X")
        
	case "linux":
		fmt.Println("GNU/Linux")
        
    case "windows": fmt.Println("Microsoft Windows")
        
	default: // freebsd, openbsd, plan9...
		fmt.Printf("%s.", os)
	} // HL
	// END OMIT
}
