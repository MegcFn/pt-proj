package handler

import (
	"fmt"
	"github.com/MegcFn/pt-proj/model"
	"github.com/MegcFn/pt-proj/util"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateApplication(c *gin.Context) {
	//validate input
	var input model.CreateApplicationInput
	var application model.Application
	db := util.GetDB()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Where("sid = ? AND grayscale = ?", input.Sid, input.Grayscale).First(&application); result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "申请已存在，请等待审批！"})
		return
	}

	application = model.Application{Username: input.Username, Sid: input.Sid, Grayscale: input.Grayscale, Reason: input.Reason}

	db.Create(&application)

	c.JSON(http.StatusCreated, gin.H{"data": application})
}

func FindApplications(c *gin.Context) {
	var applications []model.Application
	db := util.GetDB()
	db.Find(&applications)

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func FindApplicationsBySid(c *gin.Context) {
	var applications []model.Application
	db := util.GetDB()
	fmt.Println(c.Param("sid"))
	//sid, _ := strconv.ParseInt(c.Param("sid"), 10, 64)
	db.Where("sid = ?", c.Param("sid")).Find(&applications)
	//FInd方法不能得到ErrRecordNotFound
	if len(applications) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户的申请未找到!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": applications})
}

func UpdateApplication(c *gin.Context) {
	var application model.Application
	var input model.UpdateApplicationInput
	db := util.GetDB()

	if err := db.First(&application, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户申请未找到!"})
		return
	}

	fmt.Println(application)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(input)
	// 利用第三方库structs,将struct转换为map<string>interface{},解决gorm忽略空值的问题
	inputMap := structs.Map(&input)

	//判断status是否为nil，保证不审批时status不变
	if input.Flag == nil {
		db.Model(&application).Omit("flag").Updates(inputMap)
	} else {
		db.Model(&application).Updates(inputMap)
	}

	c.JSON(http.StatusCreated, gin.H{"data": application})
}

func DeleteApplication(c *gin.Context) {
	var application model.Application
	db := util.GetDB()

	if err := db.First(&application, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户申请未找到！"})
		return
	}

	db.Delete(&application)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
