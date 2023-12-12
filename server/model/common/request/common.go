package request

import "time"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page,omitempty" form:"page"`         // page number
	PageSize int    `json:"pageSize,omitempty" form:"pageSize"` // page size
    
	CreatedBefore time.Time `json:"createdBefore" form:"createdBefore"`
	CreatedAfter time.Time `json:"createdAfter" form:"createdAfter"`
    
	Order    string `json:"order" form:"order,default=DESC"`                 // ASC ||  default= DESC|

	Format   string  `json:"format" form:"format,default=Excel"`   //PDF || Excel


}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // primary key ID
}

// Uint fn
func (r *GetById) Uint() uint {
	return uint(r.ID)
}

// IdsReq struct
type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}

type IdsReqStr struct {
	Ids []string `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // role ID
}

// Empty struct
type Empty struct{}


type UserToken struct{
	ID        uint   `json:"id"`
	UserName  string `json:"userName"`
    Role      string `json:"role"`
	LastLoginTime string `json:"lastLoginTime"`
}

type LoginDTOResponse struct {
	User      UserToken       `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}
