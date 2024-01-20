package main

import "fmt"

type Bot interface {
	Hallo() string
}

type Name struct {
	Name string
}

func HelloName() Bot {
	return &Name{
		Name: "bot",
	}
}

func (n *Name) Hallo() string {
	return fmt.Sprintf("Hello %v", n.Name)
}

func main() {
	bot := HelloName()

	name := bot.Hallo()

	fmt.Println(name)
}
