// /Viết chương trình CLI cơ bản
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Định nghĩa cờ (flags)
	text := flag.String("text", "", "Text to analyze")
	flag.Parse() // Phân tích tham số

	if *text == "" {
		fmt.Println("Please provide a text using -text flag")
		return
	}

	// Đếm số từ
	words := strings.Fields(*text)
	fmt.Printf("The text contains %d words.\n", len(words))
}
