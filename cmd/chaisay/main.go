package main

import (
	"flag"
	"fmt"

	"github.com/TheChirag356/chaisay"
)

const Version = "v0.1.0"

func main() {
	str := flag.String("str", "Chaisay", "Prints the string given")
	artStyle := flag.String("art", "CHAI", "Prints your preferred art -> CHAI | COFFEE")
	version := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if *version {
		fmt.Printf("ChaiSay %s\n", Version)
		return
	}

	c := chaisay.New(*str, *artStyle)

	c.Print()
}
