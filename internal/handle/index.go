package handle

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("E:\\code\\rhim\\internal\\handle\\index.html", "E:\\code\\rhim\\front\\view\\chat\\head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
}
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("E:\\code\\rhim\\front\\view\\chat\\register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}
