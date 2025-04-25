package main

import (
	e "example/gotodo/entity"
	lstservice "example/gotodo/service/listservice"
	tservice "example/gotodo/service/taskservice"
	str "example/gotodo/storage"
	"fmt"
	"net/http"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	MsgGettingListsErrorOccurs = "error occurs getting lists: $s"
	MsgGuidNotParsed           = "guid cannot been parsed"
	MsgListJsonCannotParsed    = "error occurs parse list json"
	MsgTaskJsonCannotParsed    = "error occurs parse task json"
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

	r := gin.New()

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
	 * LISTS
	 */

	// GetAll
	r.GET("/lists", func(c *gin.Context) {
		lists, err := listService.GetAll()
		if err != nil {
			sErr := fmt.Sprintf(MsgGettingListsErrorOccurs, err)
			c.IndentedJSON(http.StatusNotFound, initMessage(sErr))
		}
		c.IndentedJSON(http.StatusOK, lists)
	})

	// GetByID
	r.GET("/lists/:id", func(c *gin.Context) {
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
	r.POST("/lists", func(c *gin.Context) {
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
	r.PUT("/lists", func(c *gin.Context) {
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
	r.DELETE("/lists/:id", func(c *gin.Context) {
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
	r.GET("/tasks/:id", func(c *gin.Context) {
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
	r.POST("/tasks", func(c *gin.Context) {
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
	r.PUT("/tasks", func(c *gin.Context) {
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
	r.DELETE("/tasks/:id", func(c *gin.Context) {
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

	r.Run("localhost:8447")
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
