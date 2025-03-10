// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package common

import (
	"chatwiki/internal/app/chatwiki/i18n"
	"chatwiki/internal/pkg/lib_web"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FmtError(c *gin.Context, msg string, params ...string) {
	data := struct{}{}
	err := errors.New(i18n.Show(GetLang(c), msg, params))
	c.String(http.StatusOK, lib_web.FmtJson(data, err))
	c.Abort()
}

func FmtErrorWithCode(c *gin.Context, code int, msg string, params ...string) {
	data := struct{}{}
	err := errors.New(i18n.Show(GetLang(c), msg, params))
	c.String(code, lib_web.FmtJsonWithCode(code, data, err))
	c.Abort()
}

func FmtOk(c *gin.Context, data interface{}) {
	c.String(http.StatusOK, lib_web.FmtJson(data, nil))
}

type response struct {
	Object    string `json:"object"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	RequestId string `json:"requestId"`
}

func FmtOpenAiErr(c *gin.Context, code int, msg string, params ...string) {
	if code == 0 {
		code = http.StatusBadRequest
	}
	err := errors.New(i18n.Show(GetLang(c), msg, params))
	c.JSON(http.StatusOK, response{
		Object:  "error",
		Message: err.Error(),
		Code:    code,
	})
	c.Abort()
}

func FmtOpenAiOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	c.Abort()
}
