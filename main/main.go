package main

import (
	"fmt"
	"net/http"
)

func HomePage() (w http.r)

func main() {
	fmt.Println("hi")
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":3000", nil)
}
