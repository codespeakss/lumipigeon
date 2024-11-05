package link

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func InitConn(urlPrefix string, username string, password string) {
	url := fmt.Sprintf("wss://%s?username=%s&password=%s", urlPrefix, username, password)
	log.Println("url: ", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("无法连接到 WebSocket:", err)
	}
	defer conn.Close()

	fmt.Printf("Connected to WebSocket server as %s\n", username)

	// 发送消息到服务端
	message := "Hello, WebSocket server!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
	fmt.Printf("Sent: %s\n", message)

	// 接收服务端回传的消息
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Error reading message:", err)
	}
	fmt.Printf("Received: %s\n", msg)

	// 等待一段时间后关闭连接
	time.Sleep(1 * time.Second)
}
