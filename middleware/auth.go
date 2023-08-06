// Package middleware
// Time : 2023/8/6 16:54
// Author : PTJ
// File : auth
// Software: GoLand
package middleware

import (
	"blog-kratos/helper"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/selector"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("Authorization")
				if token == "" {
					return nil, errors.New("no auth")
				}
				userClaims, err := helper.ParseToken(token)
				if err != nil {
					return nil, errors.New("invalid auth")
				}
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"username": []string{userClaims.Username},
				}))
			}
			return handler(ctx, req)
		}
	}
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.blog.User/Login"] = struct{}{}
	whiteList["/api.blog.User/RegisterUser"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
