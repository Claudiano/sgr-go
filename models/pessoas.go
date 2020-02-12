package models

const separator = ";"

type Person struct {
	Name       string
	Profession string
	Age        int
	Emprego
}

type Emprego struct {
	Nome  string
	Local string
}
