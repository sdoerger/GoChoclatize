package main

import (
	copy "choclatzie/pkg/copy"
	make "choclatzie/pkg/make"
	"fmt"
	"io/ioutil"
	"log"
)

const targetDir string = "gochoc"

func main() {
	// CLEAR CONSOLE
	fmt.Print("\033[H\033[2J")

	// -----------------
	// Command line Args
	// -----------------
	// cmdLnArgs := os.Args

	// Absolute path to filesdirPath + targetDir
	dirPath := "/home/stefan/Downloads/ttest"
	// TODO: CMT IN/OUT
	// dirPath := cmdLnArgs[1]

	// Limit of files into a new dir
	// TODO: CMT IN/OUT
	fileLimit := 10
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

	// Amount of dirs
	dirAmount := (filesTotal) / (fileLimit)
	// fmt.Println("dirAmount")
	// fmt.Println(dirAmount)

	// Files for last dir

	// TODO: CMNT IN
	// leftFiles := filesTotal % fileLimit
	// fmt.Println(leftFiles)

	// neededDirAmout := checkExtraLength(dirAmount, leftFiles)

	// FILES START
	s := 0
	// FILES END
	e := fileLimit
	// fmt.Println(s, e)

	// TODO: CMT IN
	make.MakeDir(dirPath, targetDir)

	// LOOP THROUGH DIR AMOUNT
	for i := 0; i < dirAmount; i++ {

		// fmt.Println("Index")
		// fmt.Println(i)

		// Copy files to dirs
		copy.CopyFiles(s, e, (dirPath), files, i, targetDir)
		s += fileLimit
		// Fle selection end
		e += fileLimit
		// fmt.Println(s, e)

		// Copy left files into last dir
	}

	// -----------------
	// Looops through files
	// -----------------
	// for _, f := range files {
	// 	// fmt.Print("File:")
	// 	fmt.Println(f.Name())
	// }

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
