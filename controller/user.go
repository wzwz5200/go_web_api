package controller

import (
	"fmt"
	"hellow/comm"
	"hellow/model"
	"log"
	"net/http"

	"hellow/response"

	"hellow/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
func WwWw(ctx *gin.Context) { 

	DB := comm.GetDB()
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	// 获取参数 名称、手机号和密码
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	fmt.Println(name,telephone,password)

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity,422,nil,"手机号必须为11位!")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity,422,nil,"密码不能少于6位!")
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity,422,nil,"用户已经存在!")
		return
	}
	// 创建用户
	// 加密用户密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError,500,nil,"密码加密失败!")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}

	DB.Create(&newUser)
	response.Success(ctx, nil,"注册成功")

}


func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	if res := db.Where("telephone = ?",telephone).First(&user);res.Error != nil {
		return false
	}
	return true
}