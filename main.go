package main

import "fmt"

type Movie struct {
	Id       string    `json:"movieid"`
	Name     string    `json:"moviename"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {

	fmt.Println("API starts")

	//Create router

	//seed testing data

	//handle routes

	//liston to the port.
}
