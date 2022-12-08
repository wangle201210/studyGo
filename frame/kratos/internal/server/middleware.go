package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
)

//MiddlewareCors 设置跨域请求头
func MiddlewareCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			//ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
			c := ctx.(http.Context)
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			//http.ResponseWriter.Header().Set()
			//if ts, ok := transport.FromServerContext(ctx); ok {
			//	if ht, ok := ts.(http.Transporter); ok {
			//		//method := ht.Request().Method
			//		//origin := ht.RequestHeader().Get("Origin")
			//		//if method == nh.MethodOptions {
			//		println("=======================")
			//		ht.RequestHeader().Set("Access-Control-Allow-Origin", "*")
			//		ht.RequestHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
			//		ht.RequestHeader().Set("Access-Control-Allow-Credentials", "true")
			//		ht.RequestHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
			//			"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
			//		//}
			//	}
			//}
			return handler(ctx, req)
		}
	}
}
