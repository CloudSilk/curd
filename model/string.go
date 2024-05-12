package model

import (
	"strings"

	"github.com/gobeam/stringy"
)

func LowerSnakeCase(s string) string {
	str := stringy.New(s)
	return str.SnakeCase().ToLower()
}

func LcFirst(s string) string {
	if s == "EID" {
		return "eid"
	}
	//TODO ID转成id
	if s == "ID" || s == "iD" {
		return "id"
	}

	str := stringy.New(s)
	return str.LcFirst()
}

func RemoveLastChar(s string) string {
	return strings.TrimRight(s, "s")
}

func CamelName(name string) string {
	array := strings.Split(name, "_")
	for i, str := range array {
		if str == "id" {
			array[i] = "id"
			continue
		}
		array[i] = strings.Title(str)
	}
	return strings.Join(array, "")
}

func CamelName2(name string) string {
	array := strings.Split(name, "_")
	for i, str := range array {
		if str == "id" {
			array[i] = "id"
			continue
		}
		if i == 0 {
			array[i] = LcFirst(str)
			continue
		}
		array[i] = strings.Title(str)
	}
	return strings.Join(array, "")
}
