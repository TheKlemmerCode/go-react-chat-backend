package main

import (
	"fmt"
	"net/http"

	"github.com/TheKlemmerCode/go-react-chat/pkg/websocket"
)

// define WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	// upgrade this connection to a WebSocket
	// connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
<<<<<<< HEAD
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

=======
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

>>>>>>> Updated the functionality of the frontend to display messages better.
	// listen indefinitely for new messages coming
	// through WebSocket connection
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	// map `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	// map `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
