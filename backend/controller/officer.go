package controller

import (
	"net/http"

	"github.com/Topzzson/SA-PROJECT/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GET /officers
// List all officers
func ListOfficer(c *gin.Context) {
	var officers []entity.Officer
	if err := entity.DB().Raw("SELECT * FROM officers").Scan(&officers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officers})
}

// GET /officers/:id
// Get officers by id
func GetOfficer(c *gin.Context) {
	var officers entity.Officer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM officers WHERE id = ?", id).Scan(&officers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officers})
}

// POST /officers
func CreateOfficer(c *gin.Context) {
	var officers entity.Officer
	if err := c.ShouldBindJSON(&officers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(officers.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	officers.Password = string(bytes)

	if err := entity.DB().Create(&officers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officers})
}

// PATCH /officers
func UpdateOfficer(c *gin.Context) {
	var officers entity.Officer
	if err := c.ShouldBindJSON(&officers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", officers.ID).First(&officers); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officers not found"})
		return
	}

	if err := entity.DB().Save(&officers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officers})
}

// DELETE /officers/:id
func DeleteOfficer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM officers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officers not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
