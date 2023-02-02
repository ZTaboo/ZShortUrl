package controller

import (
	"ZShortUrl/model"
	"ZShortUrl/service"
	"github.com/gin-gonic/gin"
)

var server = new(service.SerContext)

func GetNewUser() *User {
	return &User{}
}

func (u *User) Go(c *gin.Context) {
	path := c.Param("path")
	url, err := server.New(c).Go(path)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  err,
		})
	} else {
		c.Redirect(302, url)
	}
}

func (u *User) GenNewUrl(c *gin.Context) {
	var dataModel model.GenNewUrlModel
	res, err := server.New(c).Bind(&dataModel).GenNewUrl(dataModel)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  res,
		})
	}
}
