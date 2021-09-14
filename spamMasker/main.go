package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me something to mask")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)
	var (
		text = args[0]
		size = len(text)
		buff = make([]byte, 0, size)
		in   bool
	)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			buff = append(buff, link...)
			i += nlink
		}
		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			in = false
		}
		if in {
			c = mask
		}
		buff = append(buff, c)
	}

	fmt.Println(string(buff))

}
