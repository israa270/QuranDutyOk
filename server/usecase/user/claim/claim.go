package claim

import(
	sysReq "github.com/ebedevelopment/next-gen-tms/server/model/system/request"
	"github.com/gin-gonic/gin"
)

// GetBaseClaims
func GetBaseClaim(c *gin.Context) (*sysReq.BaseClaims, error) {
	 claims, err := GetClaims(c)
	 if err != nil{
		return nil, err
	 }
	return &claims.BaseClaims , nil
}


// GetUserInfo get of User info
func (u *ClaimUseCase) GetUserInfo(c *gin.Context) *sysReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*sysReq.CustomClaims)
		return waitUse
	}
}

