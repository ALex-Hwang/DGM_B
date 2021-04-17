package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"DGM-B/handler"
)

type GetKeyRes struct {
	Status int    `json: "status"`
	Data   string `json: "data"`
}

func RegisteRouter(e *echo.Echo) {
	GetUserAPI(e)
	GetUserInfoAPI(e)

	//设置静态路由
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	control := e.Group("/control", middleware.BasicAuth(handler.Auth))
	{
		control.GET("/getKey", func(c echo.Context) error {
			status := http.StatusOK
			room := c.QueryParam("room")
			r, err := http.Get("http://localhost:8090/control/get?room=" + room)
			if err != nil {
				panic(err)
			}
			defer func() { _ = r.Body.Close() }()

			body, _ := ioutil.ReadAll(r.Body)

			var res GetKeyRes
			json.Unmarshal(body, &res)
			fmt.Println(res)
			return c.JSON(status, map[string]string{
				"roomKey": res.Data,
			})
		})
	}

}