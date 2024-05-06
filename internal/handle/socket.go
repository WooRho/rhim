package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"rhim/internal/models"
	"rhim/middleware"
	"strconv"
	"time"
)

// 防止跨域站点伪造请求
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(c, ws)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		msg, err := middleware.Subscribe(c, middleware.PublishKey)
		if err != nil {
			fmt.Println(" MsgHandler 发送失败", err)
		}
		fmt.Println(" MsgHandler 发送成功")

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		fmt.Println(msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
func SenUserMsg(c *gin.Context) {
	middleware.Chat(c.Writer, c.Request)
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles(".\\.\\.\\front\\view\\chat\\index.html",
		".\\.\\.\\front\\view\\chat\\head.html",
		".\\.\\.\\front\\view\\chat\\foot.html",
		".\\.\\.\\front\\view\\chat\\tabmenu.html",
		".\\.\\.\\front\\view\\chat\\concat.html",
		".\\.\\.\\front\\view\\chat\\group.html",
		".\\.\\.\\front\\view\\chat\\profile.html",
		".\\.\\.\\front\\view\\chat\\main.html")
	//ind, err := template.ParseFiles(
	//	".\\.\\.\\front\\view\\chat\\main.html",
	//)
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("ToChat>>>>>>>>", user)
	ind.Execute(c.Writer, user)
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}
