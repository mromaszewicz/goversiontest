package main

// Tickle build v1-1

import (
	"log"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("could not read build info")
	}
	log.Println("Path:", bi.Main.Path)
	log.Println("Sum:", bi.Main.Sum)
	log.Println("Version:", bi.Main.Version)
}
