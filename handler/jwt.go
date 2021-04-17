package handler
import (
	"DGM-B/db"
	//"errors"
	"fmt"
	"net/http"
	"reflect"
	//"strconv"
	"time"

	//"DGM-B/models"
	//"DGM-B/middleware"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo/v4"
)

// 拓展我们要写入token的信息
type Claims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

// 认证中间件，token在userName中
func Auth(username, password string, c echo.Context) (bool, error) {
	if username == "" {
		return false, nil
	}
	extractor := basicAuthExtractor{content: username}
	token, err := request.ParseFromRequest(
		c.Request(),
		extractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secretKey"), nil
		})
	if err != nil {
		return false, err
	}
	uid := getStringFromClaims("uid", token.Claims)
	c.Set("uid",uid)

	if findByUid(uid) == false { //不存在uid
		return false, nil
	}

	return true, nil
}

// 解析 jwt token 方法
//func getUidFromClaims(key string, Rclaims jwt.Claims) string {
//	s := getStringFromClaims(key, Rclaims)
//	//value, err := strconv.Atoi(s)
//	//if err != nil {
//	//	return 0
//	//}
//	return s
//}

func getStringFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}

type basicAuthExtractor struct {
	content string
}

// basicAuthExtractor 实现 request.Extractor 接口(jwt-go下的request)
func (e basicAuthExtractor) ExtractToken(*http.Request) (string, error) {
	return e.content, nil
}

func genToken(uid string) (string, error) {
	secretKey := "secretKey"
	expiresIn := time.Duration(24 * 30)
	claims := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour.Truncate(expiresIn)).Unix(),
			Issuer:    "startdusk",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secretKey))
	return token, err
}


func findByUid(uid string) bool {
	//var us models.User
	DB_User := db.Session.DB("Live").C("user")
	cnt , err := DB_User.FindId(bson.ObjectIdHex(uid)).Count()
	//用户不存在或查询发生错误
	if err != nil || cnt==0 {
		return false
	}
	return true
}