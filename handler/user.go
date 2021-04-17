package handler

import (
	"DGM-B/db"
	"time"

	//"errors"
	"fmt"
	"net/http"

	"DGM-B/models"
	//"DGM-B/middleware"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//"github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo/v4"
)

// 从数据库中查找用户id的过程
func findUserIDFormDB(user * models.User) (string, bool) {
	var us models.User
	DB_User := db.Session.DB("Live").C("user")
	err := DB_User.Find(bson.M{"accountname":user.AccountName}).One(&us)
	//用户不存在
	if err != nil {
		return "用户不存在" , false
	}
	//密码错误
	if us.Password != user.Password{
		return "密码错误" ,false
	}
	//id转为string
	return us.ID.Hex(), true
}

// 登录模块
func Login(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	fmt.Println(u)
	if err != nil {
		return  err
	}

	id, ok := findUserIDFormDB(&u)
	if !ok {
		//return c.JSON(http.StatusUnauthorized,map[string]string{
		//	//如果错误，id为错误信息
		//	"err": id ,
		//})
		return c.String(http.StatusOK,id)
	}

	token, err := genToken(id)

	if err != nil {
		return err
	}
	return c.String(http.StatusOK,token)
	//return c.JSON(http.StatusOK, map[string]string{
	//	"token": token,
	//})
}

//注册模块
func Register(c echo.Context) error {

	var u models.User
	err := c.Bind(&u)
	fmt.Println()
	fmt.Println("user:",u)
	u.CreatedAt = time.Now()
	u.Avatar = "static/default.png"

	if err != nil {
		return  err
	}
	DB_User := db.Session.DB("Live").C("user")

	//查询是否已存在对应账户
	cnt , err := DB_User.Find(bson.M{"accountname":u.AccountName}).Count()

	if cnt != 0 {
		return c.JSON(http.StatusNotAcceptable,map[string]string{
			"err" : "账户已存在",
		})
	}

	err = DB_User.Insert(u)

	if err != nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"err":    "Fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"msg":    "Success",
	})

}

