package loanController

import (
	"net/http"
	"ujiKeterampilan/config"
	"ujiKeterampilan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Index mengambil daftar semua loan dari database dan mengembalikannya dalam format JSON.
func Index(c *gin.Context) {
	var loans []models.Loan
	config.DB.Find(&loans)
	c.JSON(http.StatusOK, gin.H{"loan": loans})
}

// Show mengambil satu data Book berdasarkan ID yang diberikan dan mengembalikannya dalam format JSON.
func Show(c *gin.Context) {
	var loan models.Loan
	id := c.Param("id")

	if err := config.DB.First(&loan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"loan": loan})
}

// Create membuat data loan baru berdasarkan data JSON yang diterima dan menyimpannya ke database.
func Create(c *gin.Context) {
	var loan models.Loan

	err := c.ShouldBindJSON(&loan)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.DB.Create(&loan)
	c.JSON(http.StatusOK, gin.H{"Loan": loan})
}

// Delete menghapus data Loan berdasarkan ID yang diberikan.
func Delete(c *gin.Context) {
	id := c.Param("id")

	var loan models.Loan
	if err := config.DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	config.DB.Delete(&loan)

	c.JSON(http.StatusOK, gin.H{"message": "loan deleted successfully"})
}
