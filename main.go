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
	} else {
		cnt := strings.Count(os.Args[1], " ")
		if cnt > 0 {
			cnt += 1
		}
		fmt.Println(cnt)
	}
}
