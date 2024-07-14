package svc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerFunc func(ctx *Context) error

type MiddleFunc func(next HandlerFunc) HandlerFunc

type OptionFunc func(*Context)

type Context struct {
	*gin.Context
	DB *gorm.DB
}

func NewContext(DB *gorm.DB) *Context {
	return &Context{
		DB: DB,
	}
}

func (e *Context) Default(code int, msg string, data any) {
	e.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func (e *Context) Success(data any) {
	e.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// Error 返回错误
func (e *Context) Error(err error) {
	e.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  err.Error(),
	})
}
