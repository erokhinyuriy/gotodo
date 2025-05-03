package main

import (
	e "example/gotodo/entity"
	"example/gotodo/middleware"
	lstservice "example/gotodo/service/listservice"
	tservice "example/gotodo/service/taskservice"
	uservice "example/gotodo/service/userservice"
	str "example/gotodo/storage"
	"fmt"
	"net/http"
	"os"
	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	MsgGettingListsErrorOccurs = "error occurs getting lists: $s"
	MsgGuidNotParsed           = "guid cannot been parsed"
	MsgListJsonCannotParsed    = "error occurs parse list json"
	MsgTaskJsonCannotParsed    = "error occurs parse task json"
	MsgUserJsonCannotParsed    = "error occurs parse user json"
	MsgErrWrongPasswordHash    = "Password is incorrect"
	MsgCannotCreateToken       = "Cannot to create a token"
)

func main() {
	fmt.Println("gotodo web api")

	// создаем хранилище
	storage, err := str.NewPostgresStorage()
	if err != nil {
		panic(err)
	}

	// сервис для работы с TdList
	listService := lstservice.New(storage)
	// сервис для работы с TdTask
	taskService := tservice.New(storage)
	// сервис для работы с User
	userService := uservice.New(storage)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies(nil)

	corsSettings := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH"},
		AllowHeaders:     []string{"Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With"},
		ExposeHeaders:    []string{"Origin"},
		AllowCredentials: true,
	})

	r.Use(corsSettings)
	r.Use(cors.Default())

	/*
	 * Auth
	 */

	r.POST("/signup", func(c *gin.Context) {
		var user e.User

		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgUserJsonCannotParsed))
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgErrWrongPasswordHash))
		}

		user.Password = string(hash)

		result, err := userService.CreateUser(&user)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}

		c.IndentedJSON(http.StatusCreated, result)
	})

	r.POST("/signin", func(c *gin.Context) {
		var userLogin e.User
		if err := c.BindJSON(&userLogin); err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgUserJsonCannotParsed))
		}
		user, err := userService.GetUser(userLogin.Email)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, initMessage(err.Error()))
			return
		}
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
		if passErr != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(passErr.Error()))
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.Id,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		tokenString, tokenErr := token.SignedString([]byte(os.Getenv("SECRET")))
		if tokenErr != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgCannotCreateToken))
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		c.IndentedJSON(http.StatusOK, tokenString)
	})

	/*
	 * LISTS
	 */

	// GetAll
	r.GET("/lists", middleware.CheckAuth, func(c *gin.Context) {
		lists, err := listService.GetAll()
		if err != nil {
			sErr := fmt.Sprintf(MsgGettingListsErrorOccurs, err)
			c.IndentedJSON(http.StatusNotFound, initMessage(sErr))
		}
		c.IndentedJSON(http.StatusOK, lists)
	})

	// GetByID
	r.GET("/lists/:id", middleware.CheckAuth, func(c *gin.Context) {
		id := c.Param("id")
		guidId, err := uuid.Parse(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgGuidNotParsed))
		}
		list, err := listService.GetByID(guidId)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, initMessage(fmt.Sprintf("%s", err)))
		}
		c.IndentedJSON(http.StatusOK, list)
	})

	// Create
	r.POST("/lists", middleware.CheckAuth, func(c *gin.Context) {
		var list e.TdList
		if err := c.BindJSON(&list); err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgListJsonCannotParsed))
		}
		guid, err := listService.Create(&list)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}

		c.IndentedJSON(http.StatusCreated, guid)
	})

	// Update
	r.PUT("/lists", middleware.CheckAuth, func(c *gin.Context) {
		var updList e.TdList
		if err := c.BindJSON(&updList); err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgListJsonCannotParsed))
		}
		result, err := listService.Update(&updList)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}
		c.IndentedJSON(http.StatusOK, result)
	})

	// Delete
	r.DELETE("/lists/:id", middleware.CheckAuth, func(c *gin.Context) {
		id := c.Param("id")
		guid, err := uuid.Parse(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgGuidNotParsed))
		}
		result, err := listService.Delete(guid)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}

		c.IndentedJSON(http.StatusOK, result)
	})

	/*
	 * TASKS
	 */

	// GetByID
	r.GET("/tasks/:id", middleware.CheckAuth, func(c *gin.Context) {
		id := c.Param("id")
		guidId, err := uuid.Parse(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgGuidNotParsed))
		}
		task, err := taskService.GetTaskByID(guidId)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, initMessage(fmt.Sprintf("%s", err)))
		}
		c.IndentedJSON(http.StatusOK, task)
	})

	// Create
	r.POST("/tasks", middleware.CheckAuth, func(c *gin.Context) {
		var task e.TdTask
		if err := c.BindJSON(&task); err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgTaskJsonCannotParsed))
		}
		guid, err := taskService.CreateTask(&task)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}

		c.IndentedJSON(http.StatusCreated, guid)
	})

	// Update
	r.PUT("/tasks", middleware.CheckAuth, func(c *gin.Context) {
		var updTask e.TdTask
		if err := c.BindJSON(&updTask); err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgTaskJsonCannotParsed))
		}
		result, err := taskService.UpdateTask(&updTask)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}
		c.IndentedJSON(http.StatusOK, result)
	})

	// Delete
	r.DELETE("/tasks/:id", middleware.CheckAuth, func(c *gin.Context) {
		id := c.Param("id")
		guid, err := uuid.Parse(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(MsgGuidNotParsed))
		}
		result, err := taskService.DeleteTask(guid)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, initMessage(fmt.Sprintf("%s", err)))
		}

		c.IndentedJSON(http.StatusOK, result)
	})

	r.Run("0.0.0.0:8447")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initMessage(msg string) map[string]interface{} {
	return gin.H{"message": msg}
}
