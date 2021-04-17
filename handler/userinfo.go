package handler
import (
	"DGM-B/db"
	"fmt"
	"io"
	"os"

	////"errors"
	//"fmt"
	"net/http"
	//"reflect"
	//"strconv"
	//"time"
	//
	"DGM-B/models"
	//"DGM-B/middleware"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo/v4"
)

//根据token获得用户信息
func GetUserInfo(c echo.Context) error {
	var u models.User
	uid := c.Get("uid").(string)
	DB_User := db.Session.DB("Live").C("user")
	selector := bson.M{"_id":bson.ObjectIdHex(uid)}
	filter := bson.M{"_id":0,"password":0}

	DB_User.Find(selector).Select(filter).One(&u)
	fmt.Println(u)
	return c.JSON(http.StatusOK,u)

}

// 用户信息修改模块
func ChangeUserInfo(c echo.Context) error {
	var u models.User
	uid := c.Get("uid").(string)

	err := c.Bind(&u)
	if err != nil{
		return err
	}
	fmt.Println(u)

	DB_User := db.Session.DB("Live").C("user")
	selector := bson.M{"_id":bson.ObjectIdHex(uid)}

	//更新传入的值
	data := bson.M{
		"$set":
			bson.M{
				"email":       u.Email,
				"phonenumber": u.PhoneNumber,
				"username":    u.Username,
				"intro":       u.Intro,
			},
	}

	err = DB_User.Update(selector,data)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
			"msg" :"OK",
		})

}

// 用户密码修改模块
func ChangePwd(c echo.Context) error {
	var u models.User
	uid := c.Get("uid").(string)

	DB_User := db.Session.DB("Live").C("user")
	selector := bson.M{"_id":bson.ObjectIdHex(uid)}

	//判断原密码是否正确
	DB_User.Find(selector).One(&u)
	if u.Password != c.FormValue("oldPWD"){
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"err" : "原密码错误",
		})
	}

	data := bson.M{
		"$set":
		bson.M{
			"password":c.FormValue("newPWD"),
		},
	}

	err := DB_User.Update(selector,data)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"msg" :"OK",
	})

}

func ChangeAvatar(c echo.Context) error {
	fmt.Println("Hello")
	uid := c.Get("uid").(string)


	// 通过echo.Contxt实例的FormFile函数获取客户端上传的单个文件
	file,err:=c.FormFile("avatar") //filename要与前端对应上
	fmt.Println(file)
	if err != nil {
		return err
	}
	// 先打开文件源
	src,err:=file.Open()
	if err != nil{
		return err
	}
	defer src.Close()

	// 下面创建保存路径文件 file.Filename 即上传文件的名字 创建upload文件夹
	dst,err := os.Create("static/"+uid+"_"+file.Filename)
	fmt.Println(dst)
	if err !=nil {
		return err
	}
	defer dst.Close()

	// 下面将源拷贝到目标文件
	if _,err = io.Copy(dst,src);err !=nil{
		return err
	}

	data := bson.M{
		"$set":
		bson.M{
			"avatar": "static/"+uid+"_"+file.Filename ,
		},
	}

	//更新头像信息
	DB_User := db.Session.DB("Live").C("user")
	selector := bson.M{"_id":bson.ObjectIdHex(uid)}
	err = DB_User.Update(selector,data)

	if err != nil {
		return err
	}

	return c.String(http.StatusOK,"头像上传成功")
}