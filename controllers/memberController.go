package controllers

import (
	"net/http"
	"ujiKeterampilan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type member struct {
	db *gorm.DB
}

func (m *member) Index(c *gin.Context) {
	var member []models.Member
	m.db.Find(&member)
	c.JSON(http.StatusOK, gin.H{"member": member})
}
