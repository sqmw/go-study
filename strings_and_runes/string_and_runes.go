package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println(utf8.RuneCountInString(s))
	for _, b := range s {
		fmt.Printf("%x\n", b)
	}

	// len(s) 仅仅统计的事字节的数量
	for i, w := 0, 0; i < len(s); i += w {
		runeVal, _w := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U\n", runeVal)
		w = _w
	}
}
