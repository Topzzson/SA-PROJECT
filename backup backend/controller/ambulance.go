package controller

import (
	"github.com/Topzzson/SA-PROJECT/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateAmbulance(c *gin.Context) {

	var ambulance entity.Ambulance

	if err := c.ShouldBindJSON(&ambulance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}

// GET
func GetAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM ambulances WHERE id = ?", id).Scan(&ambulance).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}

// GET
func ListAmbulance(c *gin.Context) {
	var ambulance []entity.Ambulance
	if err := entity.DB().Raw("SELECT * FROM ambulances").Scan(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}

// DELETE
func DeleteAmbulance(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM ambulances WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH
func UpdateAmbulance(c *gin.Context) {

	var ambulance entity.Ambulance

	if err := c.ShouldBindJSON(&ambulance); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if tx := entity.DB().Where("id = ?", ambulance.ID).First(&ambulance); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulances not found"})

		return

	}
	if err := entity.DB().Save(&ambulance).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}
