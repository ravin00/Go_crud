package main


import (
	"fmt"
	"github.com/ravin00/Go_crud/database"

)

func main(){
	database.Connect()
	fmt.Print("Successfully Connected")
}