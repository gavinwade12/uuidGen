package main

import (
	"flag"
	"fmt"
	"os"
)

const defaultVersion = 4

func main() {
	var version int
	flag.IntVar(&version, "v", 0, "the uuid version to generate")
	flag.Parse()

	if version == 0 {
		fmt.Printf("using default version: %d\n", defaultVersion)
		version = defaultVersion
	}

	app := NewApp()
	uuid, err := app.Generate(version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(uuid)
	os.Exit(0)
}
