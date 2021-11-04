package controller

import (
	"net/http"

	"github.com/Topzzson/SA-PROJECT/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GET /opertors
// List all opertors
func ListOperator(c *gin.Context) {
	var operator []entity.Operator
	if err := entity.DB().Raw("SELECT * FROM operators").Scan(&operator).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": operator})
}

// GET /opertor/:id
// Get opertor by id
func GetOperator(c *gin.Context) {
	var operator entity.Operator
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM operators WHERE id = ?", id).Scan(&operator).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": operator})
}

// POST /opertors
func CreateOperator(c *gin.Context) {
	var operator entity.Operator
	if err := c.ShouldBindJSON(&operator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(operator.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	operator.Password = string(bytes)

	if err := entity.DB().Create(&operator).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": operator})
}

// PATCH /opertors
func UpdateOperator(c *gin.Context) {
	var operator entity.Operator
	if err := c.ShouldBindJSON(&operator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", operator.ID).First(&operator); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "operator not found"})
		return
	}

	if err := entity.DB().Save(&operator).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": operator})
}

// DELETE /opertors/:id
func DeleteOperator(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM operators WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "operator not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
