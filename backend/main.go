package main

import (
	"fmt"
	"net/http"

	"github.com/3bdo-Yahya/Chatify/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
		http.Error(w, "websocket upgrade failed", http.StatusBadRequest)
        fmt.Println("upgrade error:", err)
        return
    }
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
    pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat App v0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}