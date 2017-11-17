package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// 这个例子中的map键是字符串，值是整数
	// 内置函数make创建空map
	counts := make(map[string]int)
	fmt.Println(counts)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan()  {
		//控制循环退出
		if input.Text() == "end" { break }
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

