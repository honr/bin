// Install with: $ go install bin/wsdir/wsdir
package main

import (
	"bin/wsdir"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Needs one argument")
	}
	sty := os.Args[1]
	matches, err := wsdir.Get(sty)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(matches[0])
}
