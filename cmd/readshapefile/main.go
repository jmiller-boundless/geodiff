package main

import "os"
import "github.com/jmiller-boundless/geodiff/shpfile"

func main() {
	argsWithoutProg := os.Args[1:]
	path :=argsWithoutProg[0]
	shpfile.Open(path)

}
