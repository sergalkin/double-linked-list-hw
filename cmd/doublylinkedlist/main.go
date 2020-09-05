package main

import (
	"fmt"
	"github.com/sergalkin/otus-hw-4/internal/doublylinkedlist"
	"strings"
)

func main() {
	dashes := strings.Repeat("-", 50)
	fmt.Println(dashes)

	list := doublylinkedlist.New()

	list.Push(5)
	list.Push(8)
	list.Push(13)

	fmt.Println("Длина списка:", list.Len())
	fmt.Println("Значение первого элемента в списке:", list.First().Value)
	fmt.Println("Значение второго элемента в списке:", list.First().Next().Value)
	fmt.Println("Значение последнего элемента в списке:", list.Last().Value)

	fmt.Println(dashes)

}
