package main

import (
	"fmt"
	"github.com/akutz/goswing/swingers"
)

func main() {
	var swinger swingers.Swinger
	swinger = swingers.NewSwinger()
	fmt.Printf("Hello %v\n", swinger.Name())
}
