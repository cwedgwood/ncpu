//

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
)

func cpuinfoCount() (ncpu int) {
	d, _ := ioutil.ReadFile("/proc/cpuinfo")
	for _, l := range bytes.Split(d, []byte("\n")) {
		if bytes.HasPrefix(l, []byte("processor")) {
			ncpu++
		}
	}
	return
}

func main() {
	fmt.Println("runtime.NumCPU():      ", runtime.NumCPU())
	fmt.Println("runtime.GOMAXPROCS(0): ", runtime.GOMAXPROCS(0))
	fmt.Println("/proc/cpuinfo (count): ", cpuinfoCount())
}
