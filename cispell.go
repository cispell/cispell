package main

import (
	"github.com/cispell/cispell/http"
	"github.com/cispell/cispell/config"
)

func main() {
	config.LoadEnvData()
	http.NewServer()
	http.InitServer()
	http.StartServer()
}
