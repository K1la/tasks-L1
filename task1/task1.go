package main

import "fmt"

/*
	Задание

	Дана структура Human (с произвольным набором полей и методов).

	Реализовать встраивание методов в структуре Action
	от родительской структуры Human (аналог наследования).
*/

// Дана структура Human (с произвольным набором полей,
// в данном примере только одно поле - name
// в структуру Action встраиваются поля/методы из
// родительской структуры Human (аналог наследования)
type (
	Human struct {
		name string
	}

	Action struct {
		action string
		*Human
	}
)

// Имя Human
func (h *Human) SayName() {
	fmt.Printf("My name: %s\n", h.name)
}

// действие Action
func (a *Action) SayAction() {
	fmt.Printf("My action: %s\n", a.action)
}

func main() {
	// bob с именем Bob
	bob := &Human{name: "Bob"}
	bob.SayName()

	// act с именем Action из структуры Human и
	// действие=(AAAction) из Action
	act := Action{
		Human: &Human{
			name: "Action",
		},
		action: "AAAction",
	}
	act.SayAction()
	act.SayName()
}
