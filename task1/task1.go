package main

import "fmt"

type (
	Human struct {
		name string
	}

	Action struct {
		action string
		*Human
	}
)

func (h *Human) SayName() {
	fmt.Printf("My name: %s\n", h.name)
}

func (a *Action) SayAction() {
	fmt.Printf("My action: %s\n", a.action)
}

func main() {
	bob := &Human{name: "Bob"}
	bob.SayName()

	act := Action{
		Human: &Human{
			name: "Action",
		},
		action: "AAAction",
	}
	act.SayAction()
	act.SayName()
}
