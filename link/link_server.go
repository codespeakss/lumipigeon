package link

import (
	"log"
	"net/http"
	"os"
	"time"
)

import (
	"fmt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许任何来源的请求
		return true
	},
}

func authenticate(r *http.Request) (bool, string) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// 这里可以使用你自己的认证逻辑，如数据库查询  假设我们有一个硬编码的用户名和密码
	if username == DefaultUsername && password == DefaultPassword {
		return true, username
	}
	return false, ""
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	isAuthenticated, username := authenticate(r)
	if !isAuthenticated {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Printf("New WebSocket connection established for user: %s", username)

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received from %s: %s\n", username, msg)

		// echo
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

func NewServer(port string, url string) {

	done := make(chan struct{})
	sigint := make(chan os.Signal, 1)

	go func() {
		log.Println("HandleFunc url: ", url)
		http.HandleFunc(url, handleWebSocket)
		log.Fatal(http.ListenAndServe(port, nil))
	}()

	log.Println("[NewServer] waiting select{}")
	select {
	case <-sigint:
		log.Println("Interrupted by user")
		time.Sleep(time.Millisecond * 1000)
		log.Println("Interrupted by user post")
	case <-done:
		log.Println(" chan done")
	}
}
