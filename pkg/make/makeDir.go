package make

import (
	"log"
	"os"
)

func MakeDir(p string /* dir path */, n string /* dir name */) {
	// TODO: CMT IN
	err := os.Mkdir(p+"/"+n, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(p)

}
