package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	copy "github.com/sdoerger/GoScission/pkg/copy"
	make "github.com/sdoerger/GoScission/pkg/make"
)

const targetDir string = "gosci"

func main() {
	// CLEAR CONSOLE
	fmt.Print("\033[H\033[2J")

	// -----------------
	// Command line Args
	// -----------------
	cmdLnArgs := os.Args

	// Absolute path to filesdirPath + targetDir
	dirPath := cmdLnArgs[1]
	// dirPath := "/home/stefan/Downloads/ttest"

	// Limit of files into a new dir
	fileLimit, err := strconv.Atoi(cmdLnArgs[2])
	// fileLimit := 10

	// dirMax :=

	// List files in dir from command line path
	files, err := ioutil.ReadDir(dirPath)
	// Total files
	filesTotal := len(files)

	// files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// Amount of dirs
	dirAmount := (filesTotal) / (fileLimit)

	// FILES START
	s := 0
	// FILES END
	e := fileLimit

	make.MakeDir(dirPath, targetDir)

	// -----------------
	// LOOP THROUGH DIR AMOUNT
	// -----------------
	for i := 0; i < dirAmount; i++ {

		// Copy files to dirs
		copy.CopyFiles(s, e, (dirPath), files, i, targetDir)
		s += fileLimit
		// Fle selection end
		e += fileLimit
		// fmt.Println(s, e)

	}

	// -----------------
	// LEFT FILES: Into last dir
	// -----------------
	lf := fileLimit * dirAmount
	fmt.Println(lf)

	if filesTotal-lf > 0 {
		fmt.Println("Left files")

		// DIR NAME
		dn := "dir_" + strconv.Itoa(dirAmount+1) // dir name + index

		make.MakeDir(dirPath+"/"+targetDir, dn)
		copy.CopyFiles(lf, filesTotal, (dirPath), files, dirAmount, targetDir)
	}

	fmt.Printf("Total of %d Files where moved to %d directories (%d files/dir).", filesTotal, dirAmount, fileLimit)
}

// If there are left files (not fill the last dir) add 1 to dir length
func checkExtraLength(d int, lF int) int {
	if lF >= 1 {
		return d + 1
	} else {
		return d + 0
	}
}
