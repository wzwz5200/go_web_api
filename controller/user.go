package controller

import (
	"fmt"
	"hellow/comm"
	"hellow/dto"
	"hellow/model"
	"log"
	"net/http"
	"time"

	"hellow/response"

	"hellow/utils"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Registe(ctx *gin.Context) {

	DB := comm.GetPGDB()
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	// 获取参数 名称、手机号和密码
	name := requestUser.Names
	telephone := requestUser.Telephone
	password := requestUser.Password
	fmt.Println(name, telephone, password)

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位!")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位!")
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isNameExist(DB, requestUser.Names) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在!")
		return
	}
	// 创建用户
	// 加密用户密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码加密失败!")
		return
	}
	newUser := model.User{
		Names:     name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}

	DB.Create(&newUser)
	response.Success(ctx, nil, "注册成功")

}

func isNameExist(db *gorm.DB, names string) bool {
	var user model.User
	db.Where("names = ?", names).Limit(1).Find(&user)
	if user.Names == "" {

		return false
	}
	// if res := db.Where("telephone = ?", telephone).First(&user); res.Error != nil {
	// 	return false
	// }

	return true
}

func Login(ctx *gin.Context) {
	DB := comm.GetPGDB()
	//获取数据
	//使用map获取请求参数
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	//获取参数
	name := requestUser.Names
	// telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	// fmt.Println(telephone, "手机号码长度", len(telephone))
	// if len(telephone) != 11 {
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
	// 	return
	// }
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断用户是否存在
	var user model.User
	DB.Where("names = ?", name).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := comm.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}

func Upatefile(ctx *gin.Context) {

	// 单文件
	file, _ := ctx.FormFile("./")
	log.Println(file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(file, dst)

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func Vip_Time(ctx *gin.Context) {

	DB := comm.GetPGDB()
	currentTime := time.Now()

	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	//获取参数
	name := requestUser.Names
	fast_time := model.User{
		VIP_TIME: time.Time{},
	}

	DB.Where("names = ?", name).First(&fast_time)
	if fast_time.VIP_TIME.Before(currentTime) {

		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户vip已过期")

		return
	}

	response.Success(ctx, gin.H{"VIP_time": fast_time.VIP_TIME}, "VIP未到期")
}

func Add_vip_time(ctx *gin.Context) {

	DB := comm.GetPGDB()

	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	//获取参数
	name := requestUser.Names
	fast_time := model.User{
		VIP_TIME: time.Time{},
	}

	DB.Where("names = ?", name).First(&fast_time)

	newtime := fast_time.VIP_TIME.Add(730 * time.Hour)

	fast_time.VIP_TIME = newtime

	DB.Save(&fast_time)

	response.Success(ctx, gin.H{"VIP_time": newtime}, "VIP时间已增加1个月")
}
