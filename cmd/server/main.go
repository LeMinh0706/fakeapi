package main

import (
	"github.com/LeMinh0706/fakeapi/internal/initialize"
)

func main() {
	router := initialize.InitRouter()
	router.Run()
}
