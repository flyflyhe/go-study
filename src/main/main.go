package main

import (
	"fmt"
	"services"
)

func main() {
	date := services.GetDate(-1)
	fmt.Println(date)
}
