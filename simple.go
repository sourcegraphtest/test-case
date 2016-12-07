package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

// Router is a router from gorilla/mux
type Router struct{ mux.Router }

func main() {
	fmt.Println("hello world")
}
