package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gustoliveira/speedtest-monitor/internal"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.POST("/speedtest", s.InsertSpeedtest)

	return r
}

func (s *Server) InsertSpeedtest(c *gin.Context) {
	var entry internal.SpeedtestResponse

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.db.InsertSpeedtest(entry)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, entry)
}
