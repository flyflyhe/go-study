package main

import (
	"services/trusted"
)

func main() {
	trusted.Run()
	trusted.Div2()
	trusted.InsertTrustedList(trusted.GetTrustedList(), 10000)
}
