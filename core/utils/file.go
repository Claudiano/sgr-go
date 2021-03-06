package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type File struct {
	name string
	path string
}

func NewFile(path string, nameFile string) *File {
	var file *os.File
	defer file.Close()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		check(err)
	}

	if _, err := os.Stat(path + nameFile); os.IsNotExist(err) {
		file, err = os.Create(path + nameFile)
		check(err)

	}

	return &File{path: path, name: nameFile}
}

func (f File) GetName() string {
	return f.name
}

func (f File) GetDirectoryPath() string {
	return f.path
}

func (f File) GetFilePath() string {
	return f.path + f.name
}

func (f File) WriteText(lines []string) error {
	var arq *os.File

	if _, err := os.Stat(f.GetFilePath()); os.IsNotExist(err) {
		return err
	}

	arq, err := os.OpenFile(f.GetFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer arq.Close()

	writer := bufio.NewWriter(arq)

	for _, line := range lines {
		fmt.Fprintln(writer, line)

	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return writer.Flush()
}

func (f File) WriteLine(line string) error {
	var arq *os.File

	if _, err := os.Stat(f.GetFilePath()); os.IsNotExist(err) {
		return err
	}

	arq, err := os.OpenFile(f.GetFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer arq.Close()

	writer := bufio.NewWriter(arq)
	fmt.Fprintln(writer, line)

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return writer.Flush()
}

// GetContent é uma função que le o conteudo do Files e retorna um slice the string com todas as lines do Files
func (f *File) GetContent() ([]string, error) {

	file, err := os.Open(f.GetFilePath())
	check(err)

	defer file.Close()

	// Cria um scanner que le cada line do Files
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Retorna as lines lidas e um erro se ocorrer algum erro no scanner
	return lines, scanner.Err()
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
