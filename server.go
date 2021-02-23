package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
)

func main() {

    http.HandleFunc("/", HelloHandler)
    fmt.Println("Server started at port 8080")
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%q", dump)
}