// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/consult", getConsult)
	fmt.Println("ConsultService is running on :8084")
	http.ListenAndServe(":8084", nil)
}

func getConsult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Consult Service")
}
