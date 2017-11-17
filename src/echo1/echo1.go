package main

import (
	"fmt"
	"os"
)

func main()  {
	//var s, sep string
	for i :=1; i<len(os.Args); i++ {
		fmt.Printf("%d %s", i, os.Args[i])
		fmt.Println(i)

		//s += sep + os.Args[i]
		//sep = " "
	}
	//fmt.Println(s)
}