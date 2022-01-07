package main

import (
	"fmt"
	"log"

	// "io"
	"io/ioutil"
	// "os"
)

func main() {
	// -----------------
	// Command line Args
	// -----------------
	// cmdLnArgs := os.Args

	// Absolute path to files
	dirPath := "/home/stefan/Downloads/ttest"
	// TODO: CMT IN/OUT
	// dirPath := cmdLnArgs[1]

	// Limit of files into a new dir
	// TODO: CMT IN/OUT
	fileLimit := 2
	// fileLimit := cmdLnArgs[2]

	// dirMax :=

	// List files in dir from command line path
	files, err := ioutil.ReadDir(dirPath)
	// Total files
	dirTotal := len(files)

	// files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// Amoutn of dirs
	dirAmount := (dirTotal) / (fileLimit)
	fmt.Println(dirAmount)

	// Files for last dir
	leftFiles := dirTotal % fileLimit
	fmt.Println(leftFiles)

	// -----------------
	// Looops through files
	// -----------------
	// for _, f := range files {
	// 	// fmt.Print("File:")
	// 	fmt.Println(f.Name())
	// }

	fmt.Printf("Total of %d Files \n", dirTotal)
	fmt.Printf("Limit at %v files \n", fileLimit)

}
