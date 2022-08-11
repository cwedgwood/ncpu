//

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.NumCPU():      ", runtime.NumCPU())
	fmt.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
}
