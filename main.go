package main

import (
	// "fmt"

	"github.com/IslamCHup/go-todo-json/services"
)

/*
1. читать файлы
2. вывести все тодо
3. добавить тодо
4. удалить тодо
5. отметить выполнение
*/

func main(){
	services.Add("выучить Jason")
	services.Completed("Помыть машину")
	services.Remove("выучить Jason")
	services.PrintAll()
}