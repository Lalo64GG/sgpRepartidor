package ws

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Estructura para manejar conexiones activas
type OrderConnections struct {
	Driver *websocket.Conn
	Client *websocket.Conn
}

// Mapa de conexiones por `order_id`
var orders = make(map[string]*OrderConnections)
var ordersMutex = sync.Mutex{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
// WebSocketHandler maneja las conexiones de WebSocket
func WebSocketHandler(c *gin.Context) {
	fmt.Println("Nueva conexión intentada...")

	orderID := c.Query("order_id")
	userType := c.Query("type")

	if orderID == "" || (userType != "driver" && userType != "client") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	// Actualizar a WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error al actualizar a WebSocket:", err)
		return
	}
	defer conn.Close()

	// Bloqueamos para manejar concurrencia
	ordersMutex.Lock()
	if _, exists := orders[orderID]; !exists {
		orders[orderID] = &OrderConnections{}
	}

	if userType == "driver" {
		orders[orderID].Driver = conn
	} else {
		orders[orderID].Client = conn
	}
	ordersMutex.Unlock()

	fmt.Println(userType, "conectado a la orden", orderID)

	// Escuchar mensajes del driver (ubicación)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Conexión cerrada:", err)
			break
		}

		// Si el driver envía la ubicación, la enviamos al cliente
		if userType == "driver" {
			ordersMutex.Lock()
			if orders[orderID].Client != nil {
				err = orders[orderID].Client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					fmt.Println("Error enviando ubicación al cliente:", err)
				}
			}
			ordersMutex.Unlock()
		}
	}

	// Limpiar conexiones cuando alguien se desconecta
	ordersMutex.Lock()
	if userType == "driver" {
		orders[orderID].Driver = nil
	} else {
		orders[orderID].Client = nil
	}
	ordersMutex.Unlock()
}
