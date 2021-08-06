package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)



func main() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)

	p := message.NewPrinter(language.BritishEnglish)
	p.Printf("Number format: %v.\n", 1500)

	p = message.NewPrinter(language.Greek)
	p.Printf("Number format: %v.\n", 1500)

	var a, b, c int
	fmt.Println(a, b, c)

	d := false
	fmt.Println(d)

	s := make([]int)
}

//func f(a, b int) (value int, error) { // 语法错误
//
//}

