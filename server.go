package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)

	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	if err != nil {
		fmt.Println("Error to get hostname", err)
	}

	var podInfo string
	podInfo += "<h2>Pod: "
	podInfo += hostname
	podInfo += "</h2>"

	var helloMessage string
	helloMessage += "<p>Hello, I'm "
	helloMessage += name
	helloMessage += " and have "
	helloMessage += age
	helloMessage += " old years.</p>"

	w.Write([]byte("<h1>Hello Full Cycle Rocks!!!</h1>" + podInfo + helloMessage))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("myfamily/family.txt")

	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))
}
