package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()
	echo2()
	echo3()
}

func echo1() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("echo 1: " + s)

}

func echo2() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println("echo 2: " + s)
}

func echo3() {
	fmt.Println("echo 3: " + strings.Join(os.Args[1:], " "))
}
