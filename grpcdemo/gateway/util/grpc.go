package util

import (
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

// GrpcHandlerFunc 函数是用于判断请求来源于Rpc客户端还是Restful api的请求，根据不同的请求注册不同的ServerHTTP服务，
// r.ProtoMajor=2代表着请求必须基于HTTP/2
func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}