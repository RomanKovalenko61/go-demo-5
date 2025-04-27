package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("new project")
	city := flag.String("city", "", "Город пользоввателя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)
}
