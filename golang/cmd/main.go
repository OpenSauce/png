package main

import (
	"fmt"

	"github.com/OpenSauce/png/golang/server"
)

func main() {
	fmt.Println("Starting app...")
	server := server.New("", nil)
	_ = server
}
