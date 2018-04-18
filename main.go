package main

import (
	"fmt"

	flag "github.com/ogier/pflag"
)

// command line flags
var (
	rovers    int
	gridx     int
	gridy     int
	roverpath []string
)

func main() {
	fmt.Printf("hello, world\n")
	flag.Parse()
}

func init() {
	flag.IntVarP(&rovers, "rovers", "r", 2, "The number of rovers to deploy")
	flag.IntVarP(&gridx, "gridx", "x", 5, "The X dimension of the viable terrain")
	flag.IntVarP(&gridy, "gridy", "y", 5, "They Y dimension of the viable terrain")

}
