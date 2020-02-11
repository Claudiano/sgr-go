package file

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

func (f *File) New(path string, nameFile string) {
	var file *os.File
	defer file.Close()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		check(err)
	}

	file, err = os.Create(path + nameFiles)

	f.name = nameFile
	f.path = path
}

func (f File) GetName() string {
	return f.name
}

func (f File) GetDirectoryPath() string {
	return f.path
}

func (f File) GetFilePath() string {
	return f.path
}

func (f File) WriteText(lines []string) error {
	var arq *os.File

	if _, err := os.Stat(a.GetPathFile()); os.IsNotExist(err) {
		return err
	}

	arq, err := os.OpenFile(a.Path+a.Name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer arq.Close()

	writer := bufio.NewWriter(arq)

	for _, line := range lines {
		fmt.Fprintln(writer, line)

	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return writer.Flush()
}

// GetContent é uma função que le o conteudo do Files e retorna um slice the string com todas as lines do Files
func (f *File) GetContent() ([]string, error) {

	file, err := os.Open(f.GetFilePath())
	check(err)

	defer Files.Close()

	// Cria um scanner que le cada line do Files
	var lines []string
	scanner := bufio.NewScanner(Files)
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
