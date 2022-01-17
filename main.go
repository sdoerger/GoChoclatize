package main

import (
	"fmt"
	"log"
	"os"

	// "io"
	"io/fs"
	"io/ioutil"
	"path/filepath"
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

	makeDir(dirPath, "gochoc")

	// LOOP THROUGH DIR AMOUNT
	for i := 0; i < dirAmount; i++ {

		fmt.Println("Index")
		fmt.Println(i)

		// Copy files to dirs
		copyFiles(s, e, dirPath, files, i)
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

func makeDir(p string /* dir path */, n string /* dir name */) {
	// TODO: CMT IN
	// err := os.Mkdir(p + "/" + n, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(p)

}

func copyFiles(s int /* start index */, e int /* end index */, d string /* dir path */, f []fs.FileInfo /* files */, i int /* index for dir */) {
	// s = START
	// e = END

	fmt.Println("Current Index")
	fmt.Println(i)

	dn := strings("dir_" + i) // dir name

	makeDir(d, dn)

	// Current File Selection
	cfs := f[s:e]

	for _, file := range cfs {
		// Exclude destination dir
		f := file.Name()

		if f != "gochoc" {
			// fmt.Println(f)

			// Open original file
			original, err := os.Open(d + "/" + f)
			if err != nil {
				log.Fatal(err)
			}

			defer original.Close()

			// Get file name
			fn := filepath.Base(original.Name())
			fmt.Println(fn)

			// Create new file
			new, err := os.Create(d + "/gochoc/" + fn)
			if err != nil {
				log.Fatal(err)
			}
			defer new.Close()

		}

	}

}
