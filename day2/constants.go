package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func main() {
	fmt.Printf("%d \n", 1<<10)
}
