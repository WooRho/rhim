package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"rhim/middleware"
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

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		fmt.Println(msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
