package main

import (
	"net/http"
	"op/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
