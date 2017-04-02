package main

import (
	"net/http"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Bitch has started")
	http.HandleFunc("/", indexFunc)
	http.ListenAndServe(":8080", nil)
}

func indexFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello Ben :) This is a message from a %s OS on a %s CPU", runtime.GOOS, runtime.GOARCH)
}