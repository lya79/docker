package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Hello world, the web server

	internal := os.Getenv("ENV")

	if internal == "internal" {
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			log.Println("internal/hello")
			io.WriteString(w, "internal\n")
		}

		http.HandleFunc("/hello", helloHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		return
	}

	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong\n")
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("external/hello")

		resp, err := http.Get("http://local-internal:8080/hello")
		if err != nil {
			io.WriteString(w, err.Error()+"\n")
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			io.WriteString(w, err.Error()+"\n")
			return
		}

		log.Println(string(body))
		io.WriteString(w, string(body)+"\n")
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
