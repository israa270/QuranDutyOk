package middleware

//LoadTls use https  individual middle member  exist router
//func LoadTls() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		middleware := secure.New(secure.Options{
//			SSLRedirect: true,
//			SSLHost:     "localhost:443",
//		})
//		err := middleware.Process(c.Writer, c.Request)
//		if err != nil {
//			// if appear  mistake，please not  want continue
//			fmt.Println(err)
//			return
//		}
//		// continue to   point login
//		c.Next()
//	}
//}
