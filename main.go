package main

import (
	"fmt"
	"sgr-go/core/utils"
	"sgr-go/models"
)

type x struct {
	nome string
	sexo string
}

func main() {

	p := models.Person{Name: "carlos", Profession: "marceneiro", Age: 22}
	p.Emprego.Nome = "programador"
	p.Emprego.Local = "conductor"

	fmt.Println(utils.ConvertLayoutToString(p, ";"))

	conteudo := []string{"claudiano", "sampaio"}

	a := utils.NewFile("output/", "teste 3.txt")
	a.WriteText(conteudo)
	b := utils.NewFile("output/", "teste 2.txt")

	utils.MergeFiles(a, b)
}
