package models

import (
	"fmt"
	"reflect"
	"strings"
)

const separator = ";"

type Person struct {
	Name       string
	Profession string
	Age        int
}

func (p Person) ToString() string {
	var fields []string

	v := reflect.ValueOf(p)

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, fmt.Sprintf("%v", v.Field(i)))
	}

	return strings.Join(fields, separator)

}
