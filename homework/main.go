package main

import (
	"fmt"
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
	idValue, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "id format error")
	}
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
	data := Role{}
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	Data = append(Data, data)
	c.JSON(http.StatusOK, "ok")

}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {

	data := Role{}
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	fmt.Println(data.Name)
	for i, v := range Data {
		if data.ID == v.ID {
			Data[i].Name = data.Name
			Data[i].Summary = data.Summary
			c.JSON(http.StatusOK, "ok")
			return
		}
	}

	c.JSON(http.StatusNotFound, "404 not found")
}

// 刪除資料
func Delete(c *gin.Context) {
	id := c.Param("id")
	idValue, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "id format error")
	}

	for i, items := range Data {
		if items.ID == int64(idValue) {
			Data = append(Data[:i], Data[i+1:]...)
			c.JSON(http.StatusOK, "ok")
			return
		}
	}
	c.JSON(http.StatusNotFound, "404 not found")
}
