package copy

import (
	make "choclatzie/pkg/make"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func CopyFiles(s int /* start index */, e int /* end index */, d string /* dir path */, f []fs.FileInfo /* files */, i int /* index for dir */, targetDir string) {
	// s = START
	// e = END

	fmt.Println("Current Index")
	fmt.Println(i)

	dn := "dir_" + strconv.Itoa(i+1) // dir name + index

	fmt.Println("RUNS")
	make.MakeDir(d, dn)

	// Current File Selection
	cfs := f[s:e]

	for _, file := range cfs {
		// Exclude destination dir
		f := file.Name()

		if f != targetDir {
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
			new, err := os.Create(d + targetDir + fn)
			if err != nil {
				log.Fatal(err)
			}
			defer new.Close()

		}

	}

}
