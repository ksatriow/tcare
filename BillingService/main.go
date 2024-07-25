// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/bill", getBilling)
	fmt.Println("BillingService is running on :8083")
	http.ListenAndServe(":8083", nil)
}

func getBilling(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Service Billing")
}
