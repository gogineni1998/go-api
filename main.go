package main

import (
	"fmt"

	"github.com/gogineni1998/go-api/routes"
)

func main() {
	e := routes.Routes()
	err := e.Start("0.0.0.0:8081")
	if err != nil {
		fmt.Println(err)
	}
}
