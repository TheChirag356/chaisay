package main

import (
	"flag"
)

func main() {
	str := flag.String("str", "Chaisay", "Prints the string given")
	artStyle := flag.String("art", "CHAI", "Prints your preferred art -> CHAI | COFFEE")
	flag.Parse()

	c := ChaiSay{
		Text:     *str,
		ArtStyle: *artStyle,
	}

	c.Print()
}
