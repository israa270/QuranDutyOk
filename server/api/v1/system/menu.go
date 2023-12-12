package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




type MenuApi struct {

}

type Menu struct{
ID        int       `json:"ID"`

ParentID  string    `json:"parentId"`
Path      string    `json:"path"`
Name      string    `json:"name"`
Hidden    bool      `json:"hidden"`
Component string    `json:"component"`
Sort      int       `json:"sort"`
Meta      Meta `json:"meta"`
MenuID string `json:"menuId"`
}

type Meta struct {
	ActiveName  string `json:"activeName"`
	KeepAlive   bool   `json:"keepAlive"`
	DefaultMenu bool   `json:"defaultMenu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	CloseTab    bool   `json:"closeTab"`
}


var Menus []Menu
func (s *MenuApi) GetMenu(c *gin.Context) {

	Menus = append(Menus, Menu{ 
		ID: 1,
		ParentID: "0",
		Path: "dashboard",
		Name: "dashboard",
		Hidden: false,
		Component: "view/dashboard/index.vue",
		Sort: 1,
		Meta: Meta{
			ActiveName: "",
			KeepAlive:false,
			DefaultMenu: false,
			Title: "仪表盘",
			Icon: "odometer",
			CloseTab: false,
		},

		MenuID: "1",
	
	 })
	c.JSON(http.StatusOK, gin.H{	
		"code": 0,
		"data": Menus,
		"msg": "Succeed", 
	})

}
