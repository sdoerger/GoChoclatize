package main

import (
	"fmt"
	"log"
	"os"

	// "io"
	"io/ioutil"
	// "os"
	"strings"
)

func main() {
	// -----------------
	// Command line Args
	// -----------------
	// argsWithProg := os.Args
	cmdLnArgs := os.Args[1:]
	stringed := strings.Join(cmdLnArgs, " ")

	// List files in dir from command line path
	files, err := ioutil.ReadDir(stringed)
	// files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// fmt.Print("File:")
		fmt.Println(f.Name())
	}

	fmt.Printf("%d Files \n", len(files))

}
