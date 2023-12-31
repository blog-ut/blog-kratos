package server

import (
	"blog-kratos/api/blog"
	v1 "blog-kratos/api/helloworld/v1"
	"blog-kratos/internal/conf"
	"blog-kratos/internal/service"
	"blog-kratos/middleware"
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(middleware.Auth()).Match(middleware.NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler(
		openapiv2.WithGeneratorOptions(
			generator.UseJSONNamesForFields(false),
			generator.EnumsAsInts(true)))
	srv.HandlePrefix("/q/", openAPIhandler)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	blog.RegisterUserHTTPServer(srv, service.NewUserService())
	blog.RegisterArticleHTTPServer(srv, service.NewArticleService())
	return srv
}
