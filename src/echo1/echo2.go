package main

import (
	"fmt"
	"os"
)

func main()  {
	//s, sep := "", ""
	for line, arg := range os.Args {
		fmt.Printf("%d\t%s\n", line, arg)
		//fmt.Println(arg)
		//s += sep + arg
		//sep = " "
	}
	//fmt.Println(s)
}

