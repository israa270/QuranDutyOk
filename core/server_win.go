//go:build windows
// +build windows

package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// initServer
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    30 * time.Second, // increasing timeout from 20 sec to 30 sec
		WriteTimeout:   30 * time.Second, // increasing timeout from 20 sec to 30 sec
		MaxHeaderBytes: 1 << 20,
	}
}
