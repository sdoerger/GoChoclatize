package main

import (
	"fmt"
	"log"
	"os"

	// "io"
	"io/fs"
	"io/ioutil"
	// "os"
)

func main() {
	// CLEAR CONSOLE
	fmt.Print("\033[H\033[2J")

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
	filesTotal := len(files)

	// files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// Amoutn of dirs
	dirAmount := (filesTotal) / (fileLimit)
	fmt.Println("dirAmount")
	fmt.Println(dirAmount)

	// Files for last dir
	leftFiles := filesTotal % fileLimit
	fmt.Println(leftFiles)

	// neededDirAmout := checkExtraLength(dirAmount, leftFiles)

	// FILES START
	s := 0
	// FILES END
	e := fileLimit
	fmt.Println(s, e)

	makeDir(dirPath)

	// LOOP THROUGH DIR AMOUNT
	for i := 0; i < dirAmount; i++ {

		// Copy files to dirs
		copyFiles(s, e, dirPath, files)
		s += fileLimit
		// Fle selection end
		e += fileLimit
		fmt.Println(s, e)

		// Copy left files into last dir
	}

	// -----------------
	// Looops through files
	// -----------------
	// for _, f := range files {
	// 	// fmt.Print("File:")
	// 	fmt.Println(f.Name())
	// }

	fmt.Printf("Total of %d Files \n", filesTotal)
	fmt.Printf("Limit at %v files \n", fileLimit)

}

// If there are left files (not fill the last dir) add 1 to dir length
func checkExtraLength(d int, lF int) int {
	if lF >= 1 {
		return d + 1
	} else {
		return d + 0
	}
}

func makeDir(p string /* dir path */) {
	// TODO: CMT IN
	// err := os.Mkdir(p+"/gochoc", 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(p)

}

func copyFiles(s int /* start index */, e int /* end index */, d string /* dir path */, f []fs.FileInfo /* fles */) {
	// s = START
	// e = END

	// Current File Selection
	cfs := f[s:e]

	original, err := os.Open(d + "/test_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("original")
	fmt.Println(original.Name())
	defer original.Close()

	fmt.Println("NEW SLICE")
	fmt.Println(cfs)
}
