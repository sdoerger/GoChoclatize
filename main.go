package main

import (
	"fmt"
	"log"
	"os"

	// "io"
	"io/ioutil"
	// "os"
)

func main() {
	// -----------------
	// Command line Args
	// -----------------
	cmdLnArgs := os.Args

	// Absolute path to files
	dirPath := cmdLnArgs[1]

	// Limit of files into a new dir
	fileLimit := cmdLnArgs[2]

	// dirMax :=

	// List files in dir from command line path
	files, err := ioutil.ReadDir(dirPath)
	// Total files
	dirTotal := len(files)

	// files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// fmt.Print("File:")
		fmt.Println(f.Name())
	}

	fmt.Printf("Total of %d Files \n", dirTotal)
	fmt.Printf("Limit at %v files \n", fileLimit)

}
