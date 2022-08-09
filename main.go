package main

import (
	"fmt"
	"github.com/PreBillionaire/mongoAPI/routers"
	"net/http"
)

func main() {
	r := routers.Routers()
	http.ListenAndServe(":1111",r)
	fmt.Println("Listening...")
}