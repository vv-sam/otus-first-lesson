package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Chess board size (i.e. \"8x8\"): ")
	l, w := 0, 0
	_, err := fmt.Scanf("%dx%d", &l, &w)

	if err != nil {
		panic(err)
	}

	c := 0

	sb := strings.Builder{}
	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			if c%2 == 0 {
				sb.WriteByte('#')
			} else {
				sb.WriteByte(' ')
			}
			c++
		}
		if w%2 == 0 {
			c++
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}
