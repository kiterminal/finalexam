package main

import (
	"github.com/kiterminal/finalexam/customer"
)

func main() {
	r := customer.SetupRouter()
	r.Run(":2019")
}
