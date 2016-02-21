package main

import (
	"github.com/cispell/cispell/config"
	"github.com/cispell/cispell/http"
)

func main() {
	config.LoadEnvData()
	http.NewServer()
	http.InitServer()
	http.StartServer()
}
