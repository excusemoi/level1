package task1

type Human struct {
	name string
	age  int
}

func (h *Human) printAge() {
	println(h.age)
}

func (h *Human) printName() {
	println(h.name)
}

type Action struct {
	Human //встраивание
	id string
}

func (h *Action) printAge() { //переопределение метода
	println("Age of", h.name, "is", h.age)
}

func (h *Action) printName() { //переопределение метода
	println("Name of dude with age", h.age, "is", h.name)
}

func Solve() { //демонстрационныая функцияы
	var a Action = Action{
		Human: Human{age: 12, name: "Pavel"},
		id:    "Run",
	}
	print("Action methods\n")
	a.printAge()
	a.printName()
	print("Human methods\n")
	a.Human.printAge()
	a.Human.printName()
}
