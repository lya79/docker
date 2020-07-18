package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	type User struct {
		name string
		conn *websocket.Conn
	}
	tokenMap := map[string]User{} // key: token

	// tokenMap := map[string]bool{}

	upgrader := &websocket.Upgrader{
		//如果有 cross domain 的需求，可加入這個，不檢查 cross domain
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/token" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method != "POST" {
			return
		}

		if err := r.ParseForm(); err != nil {
			log.Println("ParseForm err:", err)
			return
		}

		name := r.FormValue("name")
		password := r.FormValue("password")

		if name != "hello" || password != "world" {
			log.Println("invail name:", name, ", pwd:", password)
			return
		}

		str := ""
		str += name
		str += time.Now().String()
		token := fmt.Sprintf("%x", md5.Sum([]byte(str)))

		w.Header().Set("Content-Type", "application/json")

		type Token struct {
			Token string `json:"token"`
		}

		result := Token{
			Token: token,
		}

		tokenMap[token] = true

		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade, err:", err)
			return
		}
		defer func() {
			log.Println("disconnect")
			c.Close()
		}()

		keys, ok := r.URL.Query()["token"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}

		log.Println("tokenMap", tokenMap)

		token, ok := tokenMap[keys[0]]
		if !ok || token == false {
			log.Println("invaild token")
			return
		}

		delete(tokenMap, keys[0])

		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("receive: %s\n", msg)
			err = c.WriteMessage(mtype, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	log.Println("server start at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
