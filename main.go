package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("new project")
	city := flag.String("city", "", "Город пользоввателя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

	r := strings.NewReader("Привет! Я поток данных")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		if err == io.EOF {
			break
		}
		fmt.Printf("%q\n", block)
	}
}
