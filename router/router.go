package router

import (
	"server/internal/user"
	"server/internal/ws"
	"server/internal/message"
	

	
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, messageHandler *message.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	 
	r.POST("/api/register", userHandler.CreateUser)
	r.POST("/api/login", userHandler.Login)
	r.POST("/api/logout", userHandler.Logout)

	r.POST("/api/createRoom", wsHandler.CreateRoom)
	r.GET("/api/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/api/chat/rooms", wsHandler.GetRooms)
	r.GET("/api/chat/rooms/:roomId", wsHandler.GetClients)
	
	r.POST("/api/chat/rooms/:id/messages", messageHandler.CreateMessage)
	r.GET("/api/chat/rooms/messages", messageHandler.GetMessage)
}


func Start(addr string) error {
	return r.Run(addr)
}
