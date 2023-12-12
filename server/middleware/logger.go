package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLayout log layout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // meta customize original data
	Path      string                 // access path
	Query     string                 //query
	Body      string                 // req body data
	IP        string                 // ip address
	UserAgent string                 // agent
	Error     string                 // mistake
	Cost      time.Duration          // cost time
	Source    string                 // source
}

type Logger struct {
	// Filter User customize filter
	Filter func(c *gin.Context) bool
	// FilterKeyword Keyword filter (key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess auth point login
	AuthProcess func(c *gin.Context, layout *LogLayout)
	// log point login
	Print func(LogLayout)
	// Source serve
	Source string
}

func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = c.GetRawData()
			// use original body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      path,
			Query:     query,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		//  point login auth require  want of info
		l.AuthProcess(c, &layout)
		if l.FilterKeyword != nil {
			//  since row key/value
			l.FilterKeyword(&layout)
		}
		//  since row point login log
		l.Print(layout)
	}
}

func DefaultLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "GVA",
	}.SetLoggerMiddleware()
}
