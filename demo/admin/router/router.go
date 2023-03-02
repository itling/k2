package router

import (
	"net/http/pprof"
	"os"
	"net/http"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/middleware"
	"github.com/kingwel-xie/k2/core/logger"
	"github.com/kingwel-xie/k2/core/utils"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化
func InitRouter() {
	var r *gin.Engine
	h := common.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		common.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		logger.Fatal("not support other engine")
		os.Exit(-1)
	}

	r.Use(middleware.Sentinel()).
		Use(middleware.RequestId(utils.TrafficKey))
	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	if err != nil {
		logger.Fatalf("JWT Init Error, %s", err.Error())
	}

	r.StaticFS("/file", http.Dir("./static/file"))

	// 注册debug路由
	initDebugRouter(r)

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	// 注册业务路由
	InitBusinessRouter(r, authMiddleware)
}

func initDebugRouter(r *gin.Engine) {
	debug := r.Group("/_debug/pprof")
	{
		debug.GET("/", gin.WrapF(pprof.Index))
		debug.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		debug.GET("/profile", gin.WrapF(pprof.Profile))
		debug.POST("/symbol", gin.WrapF(pprof.Symbol))
		debug.GET("/symbol", gin.WrapF(pprof.Symbol))
		debug.GET("/trace", gin.WrapF(pprof.Trace))
		debug.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		debug.GET("/block", gin.WrapH(pprof.Handler("block")))
		debug.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		debug.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		debug.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		debug.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}

func InitBusinessRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r, authMiddleware)

	return r
}

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}

// checkRoleRouter 需要认证的路由
func checkRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v := r.Group("/api/v1")

	for _, f := range routerCheckRole {
		f(v, authMiddleware)
	}
}
