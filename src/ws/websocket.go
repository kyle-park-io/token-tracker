package ws

import (
	"net/http"
	"sync"
	"time"

	"github.com/kyle-park-io/token-tracker/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var GlobalLogChannel = make(chan string, 10)

var (
	connections = make(map[*websocket.Conn]bool)
	mu          sync.Mutex
)

func HandleWebSocket(c *gin.Context) {
	logger.Log.Infoln("Connection initiated")

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Log.Error("Failed to upgrade to WebSocket", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade WebSocket"})
		return
	}

	// Add the new connection to the connections map
	addConnection(conn)

	// Ensure the connection is removed on function exit
	defer func() {
		removeConnection(conn)
		conn.Close()
	}()

	// Set a Pong handler to verify the connection is alive
	conn.SetPongHandler(func(appData string) error {
		logger.Log.Info("Received Pong from client")
		return nil
	})

	// Start a goroutine to send periodic Ping messages to the client
	go sendPing(conn)

	// Continuously read messages from the global log channel and send them to all WebSocket clients
	for msg := range GlobalLogChannel {
		mu.Lock() // Ensure thread-safe access to the connections map
		for client := range connections {
			if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				logger.Log.Error("Failed to send message to WebSocket client", zap.Error(err))

				// Close and remove the faulty connection
				client.Close()
				delete(connections, client)
			}
		}
		mu.Unlock() // Release the lock after sending messages
	}
}

// Periodically send Ping messages to the WebSocket client to ensure the connection is active
func sendPing(conn *websocket.Conn) {
	for {
		if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			logger.Log.Error("Ping failed, closing connection", zap.Error(err))
			conn.Close()
			break
		}
		time.Sleep(30 * time.Second) // Interval between Ping messages
	}
}

// Add a WebSocket connection to the connections map
func addConnection(conn *websocket.Conn) {
	mu.Lock()
	connections[conn] = true
	mu.Unlock()
}

// Remove a WebSocket connection from the connections map
func removeConnection(conn *websocket.Conn) {
	mu.Lock()
	delete(connections, conn)
	mu.Unlock()
}

func shutdownServer() {
	logger.Log.Info("Shutting down WebSocket server...")
	closeAllConnections()
}

func closeAllConnections() {
	mu.Lock()
	defer mu.Unlock()
	for conn := range connections {
		conn.Close()              // Close the connection
		delete(connections, conn) // Remove from the map
	}
	logger.Log.Info("All connections have been closed.")
}

// Get the number of active connections
func getConnectionCount() int {
	mu.Lock()
	defer mu.Unlock()
	return len(connections)
}
