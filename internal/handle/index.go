package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	//ind, err := template.ParseFiles("index.html", "../../../front/view/chat/head.html")
	ind, err := template.ParseFiles(".\\.\\.\\index.html", ".\\.\\.\\front\\view\\chat\\head.html")
	if err != nil {
		fmt.Println("ParseFiles")
		panic(err)
	}
	err = ind.Execute(c.Writer, "")
	if err != nil {
		fmt.Println("Execute")
		panic(err)
	}
}
func ToRegister(c *gin.Context) {
	//ind, err := template.ParseFiles(".\\.\\.\\front\\view\\user\\register.html")
	ind, err := template.ParseFiles(".\\.\\.\\register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}
