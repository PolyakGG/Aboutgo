package main

import (
	"fmt"
	"net/http"
)

func helloworldpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	fmt.Println("Listening on port 9009")
	fmt.Println("i changed some shit for test123123")
	http.HandleFunc("/", helloworldpage)
	http.ListenAndServe(":9009", nil)

}
