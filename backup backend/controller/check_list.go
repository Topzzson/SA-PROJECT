package controller

import (
	"github.com/Topzzson/SA-PROJECT/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /check_lists

func CreateCheckList(c *gin.Context) {
	var owner entity.User
	var check_lists entity.CheckList
	var car_path entity.Car_path
	var path_status entity.Path_status
	var ambulance entity.Ambulance
	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร check_lists
	if err := c.ShouldBindJSON(&check_lists); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 9: ค้นหา ambulance ด้วย id
	if tx := entity.DB().Where("id = ?", check_lists.AmbulanceID).First(&ambulance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}
	// 10: ค้นหา car_path ด้วย id
	if tx := entity.DB().Where("id = ?", check_lists.Car_pathID).First(&car_path); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "car_path not found"})
		return
	}
	// 11: ค้นหา path_status ด้วย id
	if tx := entity.DB().Where("id = ?", check_lists.Path_statusID).First(&path_status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path_status not found"})
		return
	}
	// 12: ค้นหา owner ด้วย id
	if tx := entity.DB().Where("id = ?", check_lists.OwnerID).First(&owner); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path_status not found"})
		return
	}
	// 13: สร้าง checklist
	wv := entity.CheckList{

		Owner:        owner,
		Ambulance:    ambulance,
		Path_status:  path_status,
		Car_path:     car_path,
		Checked_time: check_lists.Checked_time,
	}
	// 14: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /check_list/:id
func GetCheckList(c *gin.Context) {
	var check_lists entity.CheckList
	id := c.Param("id")
	if err := entity.DB().Preload("Car_path").Preload("Ambulance").Preload("Path_status").Raw("SELECT * FROM check_lists WHERE id = ?", id).Find(&check_lists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": check_lists})

}

// GET /check_lists
func ListCheckList(c *gin.Context) {

	var check_lists []entity.CheckList

	if err := entity.DB().Preload("Ambulance").Preload("Car_path").Preload("Path_status").Raw("SELECT * FROM check_lists").Find(&check_lists).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	c.JSON(http.StatusOK, gin.H{"data": check_lists})
}

// DELETE /check_lists/:id
func DeleteCheckList(c *gin.Context) {
	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM check_lists WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "check_lists not found"})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /check_lists
func UpdateCheckList(c *gin.Context) {

	var check_lists entity.CheckList

	if err := c.ShouldBindJSON(&check_lists); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if tx := entity.DB().Where("id = ?", check_lists.ID).First(&check_lists); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "check_lists not found"})
		return
	}

	if err := entity.DB().Save(&check_lists).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": check_lists})
}
