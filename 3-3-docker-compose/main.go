package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Hello world, the web server

	internal := os.Getenv("ENV")

	if internal == "internal" {
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			config := fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
				"root",
				"qwe123",
				"db:3306",
				"hello",
				true,
				"Local",
			)
			db, err := gorm.Open("mysql", config)
			if err != nil {
				io.WriteString(w, "open:"+err.Error()+"\n")
				return
			}
			defer db.Close()

			db.LogMode(true)

			type User struct {
				ID   int64  `gorm:"id"`
				Name string `gorm:"name"`
			}

			arr := []User{}
			err = db.Table("user").Scan(&arr).Error
			if err != nil {
				io.WriteString(w, "scan:"+err.Error()+"\n")
				return
			}
			str := ""
			for _, s := range arr {
				str += s.Name + ","
			}
			io.WriteString(w, str+"\n")
		}

		http.HandleFunc("/hello", helloHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
		return
	}

	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "pong\n")
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
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

		io.WriteString(w, string(body))
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
