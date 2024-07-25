// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/book", getBook)
	fmt.Println("BookService is running on :8082")
	http.ListenAndServe(":8082", nil)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Book Service")
}
