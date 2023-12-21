package middleware

//// NeedInit point login cross domain please request,support options access
//func NeedInit() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if global.GvaDB == nil {
//			response.OkWithDetailed(gin.H{
//				"needInit": true,
//			}, "before to initialization database", http.StatusOK, "success", c)
//			c.Abort()
//		} else {
//			c.Next()
//		}
//	}
//}
