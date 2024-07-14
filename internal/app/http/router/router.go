package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"etoc-service/internal/app/db"
	"etoc-service/internal/app/http/svc"
)

type OptionFunc func(r *Router)

func WithSqlite() OptionFunc {
	return func(r *Router) {
		r.db = db.SqliteDB()
	}
}

type Router struct {
	engine     *gin.Engine
	middleFunc []svc.MiddleFunc
	basePath   string
	db         *gorm.DB
}

func NewRouter(engine *gin.Engine, options ...OptionFunc) *Router {
	r := &Router{
		engine:     engine,
		middleFunc: make([]svc.MiddleFunc, 0),
		basePath:   "",
	}
	for _, opt := range options {
		opt(r)
	}
	return r
}

func (r *Router) routerHandler(handler svc.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := svc.NewContext(r.db)
		context.Context = ctx
		for _, middleFunc := range r.middleFunc {
			handler = middleFunc(handler)
		}
		err := handler(context)
		if err == nil {
			return
		}
		context.Error(err)
		logrus.Error(err)
	}
}

func (r *Router) Group(relativePath string, middleFuncs ...svc.MiddleFunc) *Router {
	middleFunc := append(r.middleFunc, middleFuncs...)
	return &Router{
		engine:     r.engine,
		middleFunc: middleFunc,
		basePath:   r.basePath + relativePath,
	}
}

func (r *Router) POST(path string, handler svc.HandlerFunc) {
	r.engine.POST(r.basePath+path, r.routerHandler(handler))
}

func (r *Router) GET(path string, handler svc.HandlerFunc) {
	r.engine.GET(r.basePath+path, r.routerHandler(handler))
}

func (r *Router) Static(path string, root string) {
	r.engine.Static(path, root)
}

func (r *Router) RegisterHandler(handler func(r *Router)) {
	handler(r)
}
