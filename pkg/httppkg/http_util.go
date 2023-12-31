package httppkg

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/go-gin-template/pkg/stringspkg"
)

func GetURLParam(ctx *gin.Context, key string) string {
	return ctx.Param(key)
}

func GetQueryParam(ctx *gin.Context, key string) string {
	return ctx.Query(key) // shortcut for c.Request.URL.Query().Get("lastname")
}

func ParseAuthToken(ctx *gin.Context) ([]string, error) {
	h := ctx.Request.Header["Authorization"]
	if h == nil && len(h) == 0 {
		return nil, errors.New("auth token is missing")
	}
	tkHeader := h[0]
	if tkHeader == "" {
		// Token is missing
		return nil, errors.New("auth token is missing")
	}
	splits := stringspkg.Split(tkHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		return nil, errors.New("token format is invalid")
	}
	return splits, nil
}

func GetFile(r *http.Request, k string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(k)
}
