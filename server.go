package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("Error to get hostname", err)
	}

	var podInfo string
	podInfo += "<h2>Pod: "
	podInfo += hostname
	podInfo += "</h2>"

	w.Write([]byte("<h1>Hello Full Cycle Rocks!!!</h1>" + podInfo))
}
