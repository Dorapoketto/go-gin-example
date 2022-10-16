package api

import (
	"fmt"
	"github.com/Dorapoketto/go-gin-example/models"
	"github.com/Dorapoketto/go-gin-example/pkg/e"
	"github.com/Dorapoketto/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type auth struct {
	Username string `validate:"required,max=50"`
	Password string `validate:"required,max=50"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	validate := validator.New()
	a := auth{Username: username, Password: password}
	err := validate.Struct(a)
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	isExist := models.CheckAuth(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
