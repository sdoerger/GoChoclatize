package make

import (
	"fmt"
	"log"
	"os"
)

func MakeDir(p string /* dir path */, n string /* dir name */) {
	// CHECK: If dir exists
	_, err := os.Stat(p + "/" + n)
	if os.IsNotExist(err) {

		// CREATE DIR: With rw rights
		errr := os.Mkdir(p+"/"+n, 0755)
		if errr != nil {
			fmt.Println("Error from MakeDir")
			log.Fatal(errr)
		}
	}
}
