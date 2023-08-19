package memberController

import (
	"net/http"
	"ujiKeterampilan/config"
	"ujiKeterampilan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Index mengambil daftar semua member dari database dan mengembalikannya dalam format JSON.
func Index(c *gin.Context) {
	var members []models.Member
	config.DB.Find(&members)
	c.JSON(http.StatusOK, gin.H{"member": members})
}

// Show mengambil satu data Book berdasarkan ID yang diberikan dan mengembalikannya dalam format JSON.
func Show(c *gin.Context) {
	var member models.Member
	id := c.Param("id")

	if err := config.DB.First(&member, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"member": member})
}

// Create membuat data Member baru berdasarkan data JSON yang diterima dan menyimpannya ke database.
func Create(c *gin.Context) {
	var member models.Member

	err := c.ShouldBindJSON(&member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	config.DB.Create(&member)
	c.JSON(http.StatusOK, gin.H{"member": member})
}

// // Update mengupdate data Event berdasarkan ID yang diberikan dengan data JSON yang diterima.
// func Update(c *gin.Context) {
// 	var event models.Event
// 	id := c.Param("id")

// 	if err := c.ShouldBindJSON(&event); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	if models.DB.Model(&event).Where("id = ?", id).Updates(&event).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate Event"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
// }

// Delete menghapus data Book berdasarkan ID yang diberikan.
func Delete(c *gin.Context) {
	id := c.Param("id")

	var member models.Member
	if err := config.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		return
	}

	config.DB.Delete(&member)

	c.JSON(http.StatusOK, gin.H{"message": "member deleted successfully"})
}
