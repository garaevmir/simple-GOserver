package main

func main() {
	name := "Amir"
	pass := "123456789"

	AddNewUser(name, pass)
	AddInfo(name, "Some information")

	name = "Ivan"
	AddNewUser(name, pass)
	AddInfo(name, "Another interesting information")

	name = "John"
	AddNewUser(name, pass)
	AddInfo(name, "Absolutely vital info")
}
