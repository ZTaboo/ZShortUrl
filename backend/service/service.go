package service

import (
	"ZShortUrl/client"
	"ZShortUrl/config"
	"ZShortUrl/model"
	"ZShortUrl/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

type SerContext struct {
	c *gin.Context
}

func (s *SerContext) New(c *gin.Context) *SerContext {
	s.c = c
	return s
}

func (s *SerContext) Bind(obj any) *SerContext {
	if s.c == nil {
		fmt.Println("lack *gin.Context")
		return s
	}
	err := s.c.ShouldBindJSON(obj)
	if err != nil {
		fmt.Println("解析参数错误：", err)
		name, _, _, _ := runtime.Caller(1)
		fmt.Println("调用者为：", runtime.FuncForPC(name).Name())
	}
	return s
}

func (s *SerContext) Go(path string) (string, error) {
	fmt.Println("path:", path)
	url, err := client.Cache.Get(path)
	fmt.Println("url:", url)
	if err != nil {
		return "", err
	}
	return string(url), nil
}

func (s *SerContext) GenNewUrl(dataModel model.GenNewUrlModel) (string, error) {
	//生成随机数
	res := utils.RandSeq(5)
	fmt.Println("dataModel", dataModel)
	err := client.Cache.Set(res, []byte(dataModel.Url))
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	return config.Conf.Base.Url + "/" + res, nil
}
