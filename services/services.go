package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IslamCHup/go-todo-json/structure"
)

/*
1. читать файлы
2. вывести все тодо
3. добавить тодо
4. удалить тодо
5. отметить выполнение
*/


func ReadTodo() ([]structure.Todo, error) {
	cont, err := os.ReadFile("todo.json")
	if err != nil{
		return nil, err
	}
	
	var t []structure.Todo

	errorReadJS := json.Unmarshal(cont, &t)
	if errorReadJS != nil{
		return nil, errorReadJS
	}



	return t, nil
}

func WriteTodo(content []structure.Todo)error{
	jsdata, errorJS := json.MarshalIndent(content, "", "  ")

	if errorJS != nil{
		return fmt.Errorf("%s", errorJS)
	}

	errorWrite := os.WriteFile("todo.json", jsdata, 0644)

	if errorWrite != nil{
		return fmt.Errorf("%s", errorWrite)
	}
	return 	nil
}
// todo >> marshal >> os.wrtie
// content >> unmarshal >> add todo >> marshal >> os.write
func Add(text string)error{

	content, err := ReadTodo()

	if err != nil{
		return fmt.Errorf("%s", err)
	}


	t := structure.Todo{Text: text, Complete: false}
	content = append(content, t)

	errorsWrite := WriteTodo(content)

	if errorsWrite != nil{
		return fmt.Errorf("%s", errorsWrite)
	}

	return nil
}

func Remove(title string)error{
	cont, _:= ReadTodo()

	resSlice := []structure.Todo{}
	
	for _, val := range cont{
		if val.Text != title{
			resSlice = append(resSlice, val)
		}
	} 

	cont = resSlice

	err := WriteTodo(cont)
	if err != nil{
		return fmt.Errorf("%s",err)
	}
	return nil
}

func Completed(title string){
	cont, _ := ReadTodo()
	for i := range cont{
		if cont[i].Text == title {
			cont[i].Complete = !cont[i].Complete
		}
	}

	WriteTodo(cont)
}


func PrintAll()error{
	content, err := ReadTodo()

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	complete := ""
	for _, v := range content{
		if v.Complete{
			complete = "[v]"
		} else{
			complete = "[ ]"
		}
		fmt.Printf("%s %s \n", v.Text, complete)
	}
	return nil
}




