package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/camelcase"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/ebedevelopment/next-gen-tms/server/model/system"
	sysResp "github.com/ebedevelopment/next-gen-tms/server/model/system/response"
	claimCase "github.com/ebedevelopment/next-gen-tms/server/usecase/user/claim"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.GvaLog.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, _ := claimCase.GetClaims(c)
		if claims.ID != 0 {
			userId = int(claims.ID)
		} else {
			id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}

		record := system.SysOperationRecord{
			IP:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
	
		}

		//handle Name of operation
		operationName := strings.Split(c.Request.URL.Path, "/")
		name := strings.Split(operationName[2], "?")
		split := camelcase.Split(name[0])
		record.ActionType = split[0]
		record.Name = strings.Join(split, " ")
		record.Name = strings.ToLower(record.Name)

		// exist 
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}

		// if strings.Index(c.GetHeader("Content-Type"), "multipart/form-data") > -1 {
		// 	if len(record.Body) > 1024 {
		// 		// truncate
		// 		newBody := respPool.Get().([]byte)
		// 		copy(newBody, record.Body)
		// 		record.Body = string(newBody)
		// 		defer respPool.Put(newBody[:0])
		// 	}
		// }

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		// Read the response body as a byte slice

		record.Resp = writer.body.String()
		// for import excel file
		values := c.Request.Header.Values("content-type")
		if len(values) > 0 && strings.Contains(values[0], "multipart/form-data") {

			record.Resp = "{'message':'file success'}"
		}

		//  for export excel
		bodyStr := writer.body.String()
		if strings.Contains(bodyStr, "workbook.xml") || strings.Contains(bodyStr, "xl/worksheets") || strings.Contains(bodyStr, "import success") {
			record.Resp = "{'message':'file success'}"
		}
		// end

		latency := time.Since(now)
		// record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		if record.Status == 200 {
			record.StatusOp = "Success"
		} else {
			record.StatusOp = "Fail"
		}

		record.Latency = latency
		record.Resp = writer.body.String()

		// if strings.Index(c.Writer.Header().Get("Pragma"), "public") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Expires"), "0") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Content-Type"), "application/force-download") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Content-Type"), "application/json") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Content-Type"), "application/download") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Content-Disposition"), "attachment") > -1 ||
		// 	strings.Index(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") > -1 {
		// 	if len(record.Resp) > 1024 {
		//
		// 		newBody := respPool.Get().([]byte)
		// 		copy(newBody, record.Resp)
		// 		record.Body = string(newBody)
		// 		defer respPool.Put(newBody[:0])
		// 	}
		// }

		if len(record.Body) > 65535 {
			record.Body = record.Body[:65534]
		}

		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			if !strings.Contains(err.Error(), "Error 1366: Incorrect string value:") || !strings.Contains(err.Error(), "Error 1406: Data too long") {
				global.GvaLog.Error("create operation record error:", zap.Error(err))
			}
		}

		op := sysResp.RecordOperation{
			Method:   record.Method,
			Ip:       record.IP,
			Path:     record.Path,
			Status:   record.Status,
			StatusOp: record.StatusOp,
			Latency:  fmt.Sprintf("%v", time.Duration(record.Latency.Microseconds())),
			Agent:    record.Agent,
			Body:     record.Body,
			Resp:     record.Resp,
			// ErrorMessage: record.ErrorMessage,
			UserName:   claims.Username,
			// UserEmail:  claims.Email,
			Name:       record.Name,
			ActionType: record.ActionType,


			// AuthorityId: claims.AuthorityId,
		}
		// store operation in log file
		global.GvaLog.Info(fmt.Sprintf("%#v", op))
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
