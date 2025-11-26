package handler

import (
	"net/http"

	"gin-demo/service"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	username := c.Query("username")

	user, msg, err := service.Login(username)
	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		case service.ErrUserNotActive:
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		case service.ErrUnknownRole:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"user":    user,
	})
}
