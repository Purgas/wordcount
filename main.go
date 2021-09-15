package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//var str = flag.String("str", "", "set string")
	//flag.Parse()
	if len(os.Args) == 1 {
		fmt.Println(0)
	} else if len(os.Args[1]) == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(strings.Count(os.Args[1], " ") + 1)
	}
}
