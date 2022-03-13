package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var formatted, err = httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Println(string(formatted))

	w.Write(formatted)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
