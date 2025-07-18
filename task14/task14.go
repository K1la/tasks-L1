package main

import (
	"fmt"
	"reflect"
)

/*
	Определение типа переменной в runtime

	Разработать программу, которая в runtime способна
	определить тип переменной, переданной в неё
	(на вход подаётся interface{}).
	Типы, которые нужно распознавать:
	int, string, bool, chan (канал).

    -------------------------------------------------------
	Для определения типа канала используется либа "reflect"
*/

func main() {
	detect(1)
	detect(1.2)
	ch := make(chan int)
	detect(ch)
	chstr := make(chan string)
	detect(chstr)
	chst := make(chan struct{}, 1)
	detect(chst)
}

func detect(v interface{}) {
	switch x := v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Chan {
			fmt.Printf("chan of %s\n", rv.Type().Elem())
		} else {
			fmt.Println("не поддерживаемый:", x)
		}
	}
}
