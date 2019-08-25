package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type Ws int

func (Ws) Server(c *gin.Context) {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	online(c.Writer, c.Request, uint(id))
}

func (Ws) Offline(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	delete(Hub, uint(id))
	fmt.Println("下线")
}

func (Ws) Dispatch(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	conn, has := Hub[uint(id)]
	if has {
		if err := conn.WriteMessage(2, []byte("更新你的列表")); err != nil {
			log.Fatal(err)
		}
	}
}

// Register register a group of router to root router
func (w Ws) Register(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/ws/:id", w.Server)
	st.DELETE("/ws/:id", w.Offline)
	st.PUT("/ws/:id", w.Dispatch)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Hub map[uint]*websocket.Conn

func online(w http.ResponseWriter, r *http.Request, id uint) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 存储
	Hub[id] = conn
	fmt.Println("上线")
}

func init() {
	Hub = make(map[uint]*websocket.Conn, 0)
}
