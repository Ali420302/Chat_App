package message

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
	
}




func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateMessage(c *gin.Context) {
	var u CreateMessageReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    

	res, err := h.Service.CreateMessage(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetMessage(c *gin.Context) {
	u := CreateMessageReq{}
	
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    

	// Query the messages by roomId
query := "SELECT * FROM messages WHERE RoomID = $1"
rows := query


	c.JSON(http.StatusOK, rows)
}