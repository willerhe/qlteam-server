package rest

import "github.com/gin-gonic/gin"

type Story int

func (s *Story) get(c *gin.Context) {

}

func (s *Story) RegisterStory(r *gin.RouterGroup) {
	st := r.Group("")
	st.GET("/story/:id", s.get)

}
