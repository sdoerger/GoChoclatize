package make

import (
	"fmt"
	"log"
	"os"
)

func MakeDir(p string /* dir path */, n string /* dir name */) {
	// TODO: CMT IN

	fmt.Println(p + "/" + n)

	// CHECK: If dir exists
	folderInfo, err := os.Stat(p + "/" + n)
	if os.IsNotExist(err) {

		errr := os.Mkdir(p+"/"+n, 0755)
		if errr != nil {
			fmt.Println("Error from MakeDir")
			log.Fatal(errr)
		}
		fmt.Printf("Dir %s was created).", n)
		// fmt.Println(p)

	}
	log.Println(folderInfo)

}
