package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut65/team11/entity"
)

// POST Type
func CreateType(c *gin.Context) {
	var types entity.Type
	if err := c.ShouldBindJSON(&types); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&types).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": types})
}

func GetListType(c *gin.Context) {
	var types []entity.Type
	if err := entity.DB().Raw("SELECT * FROM types").Scan(&types).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": types})
}

// POST Windows
func CreateWindows(c *gin.Context) {
	var windows entity.Windows
	if err := c.ShouldBindJSON(&windows); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&windows).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": windows})
}

func GetListWindows(c *gin.Context) {
	var windows []entity.Windows
	if err := entity.DB().Raw("SELECT * FROM windows").Scan(&windows).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": windows})
}

// Main Table Device
func CreateDevice(c *gin.Context) {

	var customer entity.Customer
	var types entity.Type
	var windows entity.Windows
	var device entity.Device

	// ผลลัพธ์ที่ได้จะถูก bind เข้าตัวแปร Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา Type  ด้วย id
	if tx := entity.DB().Where("id = ?", device.TypeID).First(&types); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type  not found"})
		return
	}

	// ค้นหา Windows ด้วย id
	if tx := entity.DB().Where("id = ?", device.WindowsID).First(&windows); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Windows  not found"})
		return
	}

	// 13: ค้นหา Customer ด้วย id
	if tx := entity.DB().Where("id = ?", device.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	dv := entity.Device{
		CPU:        device.CPU,        // ตั้งค่าฟิลด์ CPU
		Monitor:    device.Monitor,    // ตั้งค่าฟิลด์ Monitor
		GPU:        device.GPU,        // ตั้งค่าฟิลด์ GPU
		RAM:        device.RAM,        // ตั้งค่าฟิลด์ RAM
		Harddisk:   device.Harddisk,	// ตั้งค่าฟิลด์ Harddisk
		Problem:	device.Problem ,	//ตั้งค่าฟิลด์ Problem
		CustomerID: device.CustomerID, // โยงความสัมพันธ์กับ Entity Customer
		TypeID:     device.TypeID,     // โยงความสัมพันธ์กับ Entity Type
		WindowsID:  device.WindowsID,  // โยงความสัมพันธ์กับ Entity Windows
		Save_Time:  device.Save_Time,  // ตั้งค่าฟิลด์ Save_Time
	}

	if err := entity.DB().Create(&dv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dv})
}

// GET /Device
func GetListDevice(c *gin.Context) {
	var devices []entity.Device
	if err := entity.DB().Preload("Customer").Preload("Type").Preload("Windows").Find(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// GET /Device:id
func GetDevice(c *gin.Context) {
	var device entity.Device
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM devices WHERE id = ?", id).Scan(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}

// PATCH /Device
func UpdateDevice(c *gin.Context) {
	var device entity.Device

	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Model(device).Where("id = ?", device.ID).Updates(map[string]interface{}{"CPU": device.CPU, "Monitor": device.Monitor, "GPU": device.GPU, "RAM": device.RAM, "Harddisk": device.Harddisk, "Problem": device.Problem, "CustomerID": device.CustomerID, "TypeID": device.TypeID, "WindowsID": device.WindowsID, "Save_Time": device.Save_Time}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": device})
}

// DELETE /Device
func DeleteDevice(c *gin.Context) {
	var device entity.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Exec("DELETE FROM devices WHERE id = ?", device.ID); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "DELETE SUCCEED!!"})
}
