package router

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	})
	r.GET("/ping", h.ping)
	r.POST("/registration", h.ident)
	r.POST("/login", h.login)
	r.GET("/questions", h.getQuestions)
	r.OPTIONS("/questions/:id", h.deleteQuestions)
	r.POST("/question/update", h.updateQuestions)
	r.POST("/answers", h.getAnswers)
	r.GET("/user/answers", h.getUserAnswers)
	return r
}

func (h *Handlers) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (h *Handlers) ident(c *gin.Context) {

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 json"})
		return
	}
	usr, err := h.Srv.CheckNewUser(&user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
		} else {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
	}
	if usr != nil && usr.Id != 0 {
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

	c.JSON(http.StatusOK, gin.H{"role": usr.Role, "user_id": usr.Id})
}

func (h *Handlers) getQuestions(c *gin.Context) {

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
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	var newQuest []models.Questions
	err := c.ShouldBindJSON(&newQuest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	questions, err := h.Srv.UpdateQuestions(newQuest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func (h *Handlers) deleteQuestions(c *gin.Context) {
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

func (h *Handlers) getAnswers(c *gin.Context) {
	incomingData := make(map[string]interface{})

	err := c.ShouldBindJSON(&incomingData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	answers, err := h.Srv.AddAnswers(incomingData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": answers})
}

func (h *Handlers) getUserAnswers(c *gin.Context) {

	answers, err := h.Srv.GetAnswers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": answers})
}
