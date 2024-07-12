package main

import (
	"github.com/gogineni1998/go-api/routes"
)

func main() {
	e := routes.Routes()
	e.Start("0.0.0.0:8081")
}
