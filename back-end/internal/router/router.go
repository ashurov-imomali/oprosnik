package router

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/service"
	"main/models"
	"net/http"
	"strconv"
)

type Handlers struct {
	Srv *service.Service
}

func InitHandler(srv *service.Service) *Handlers {
	return &Handlers{Srv: srv}
}

func GetHandlers(h *Handlers) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", h.ping)
	r.POST("/registration", h.ident)
	r.POST("/login", h.login)
	r.GET("/questions", h.getQuestions)
	r.DELETE("/questions/:id", h.deleteQuestions)
	r.PUT("/questions", h.updateQuestions)
	return r
}

func (h *Handlers) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (h *Handlers) ident(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 json"})
		return
	}
	usr, err := h.Srv.CheckNewUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if usr.Id != 0 {
		log.Println("current user")
		c.JSON(http.StatusBadRequest, gin.H{"message": "current user"})
		return
	}
	newUser, err := h.Srv.CreateUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}

func (h *Handlers) login(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	usr, err := h.Srv.CheckNewUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if ok := h.Srv.VerifyPass(&user, usr.Password); !ok {
		log.Println("incorrect pass", user.Password, fmt.Sprintf("%x", md5.Sum([]byte(user.Password))))
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": usr.Role})
}

func (h *Handlers) getQuestions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	questions, err := h.Srv.GetQuestions()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func (h *Handlers) updateQuestions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	var newQuest models.Questions
	err := c.ShouldBindJSON(&newQuest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	questions, err := h.Srv.UpdateQuestions(&newQuest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func (h *Handlers) deleteQuestions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = h.Srv.DeleteQuestions(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
}
