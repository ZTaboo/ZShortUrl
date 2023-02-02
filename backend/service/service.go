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
	c      *gin.Context
	param  string
	params []string
}

func (s *SerContext) New(c *gin.Context) *SerContext {
	s.c = c
	return s
}
func (s *SerContext) Param(key string, keys ...string) *SerContext {
	switch len(keys) {
	case 0:
		s.param = s.c.Param(key)
		break
	default:
		s.param = s.c.Param(key)
		for _, item := range keys {
			s.params = append(s.params, s.c.Param(item))
		}
	}
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

func (s *SerContext) Go() (string, error) {
	url, err := client.Cache.Get(s.param)
	if err != nil {
		return "", err
	}
	return string(url), nil
}

func (s *SerContext) GenNewUrl(dataModel model.GenNewUrlModel) (string, error) {
	//生成随机数
	res := utils.RandSeq(5)
	err := client.Cache.Set(res, []byte(dataModel.Url))
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	return config.Conf.Base.Url + "/" + res, nil
}
