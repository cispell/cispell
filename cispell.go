package main

import (
	"github.com/cispell/cispell/http"
)

func main() {
	http.NewServer()
	http.InitServer()
	http.StartServer()
}
