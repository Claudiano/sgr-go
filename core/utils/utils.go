package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func ConvertLayoutToString(s interface{}, separator string) string {
	var fields []string

	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, fmt.Sprintf("%v", v.Field(i)))
	}

	return strings.Join(fields, separator)

}

func MergeFiles(source *File, from *File) {

	fileIn, err := os.Open(source.GetFilePath())
	if err != nil {
		log.Fatalln("Erro ao abrir arquivo %v", err)
	}

	defer fileIn.Close()

	fileOut, err := os.OpenFile(from.GetFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Erro ao abrir arquivo %v", err)
	}

	defer fileOut.Close()

	_, err = io.Copy(fileOut, fileIn)
	if err != nil {
		log.Fatalln("Erro ao realizar merge dos arquivos, error: %v", err)
	}

	log.Println("Merge finalizado")

}
