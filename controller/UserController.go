package controller

import (
	"ginessential/common"
	"ginessential/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = RandomString(10)
	}

	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// 返回结果

	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
func RandomString(n int) string {
	var letters = []byte("ssssssssssss")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
