package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/Role", Get)

	router.GET("/Role/:id", GetRoleById)

	router.POST("/Role", Post)

	router.PUT("/Role/:id", Put)

	router.DELETE("/Role/:id", Delete)

	router.Run(":8080")

}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetRoleById(c *gin.Context) {
	id := c.Param("id")
	idValue, _ := strconv.Atoi(id)
	for _, items := range Data {
		if items.ID == int64(idValue) {
			c.JSON(http.StatusOK, items)
		} else {
			c.JSON(http.StatusNotFound, "no data")
		}
	}
}

// 新增資料
func Post(c *gin.Context) {

}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {

}

// 刪除資料
func Delete(c *gin.Context) {

}
