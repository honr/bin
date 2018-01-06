// Install with: $ go install bin/tm
package main

import (
	"bin/wsdir"
	"fmt"
	"log"
	"os"
	// "cmd"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("TODO: list existing sessions")
	}
	matches, err := wsdir.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.ExpandEnv(matches[0]))
}
