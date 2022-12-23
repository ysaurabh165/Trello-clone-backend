package router

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"time"
	controllers "github.com/ysaurabh165/Trello-clone-backend/internal/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jacobstr/confer"
)

func Setup(configuration *confer.Config) *gin.Engine {
	router := gin.New()
	filePath := configuration.GetString("logfile")
	f,_ := os.Create(filePath)
	fmt.Println("filePath======="+filePath)
	gin.DefaultWriter = io.MultiWriter(f)

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/login", controllers.GetLogin)

	router.GET("/boardContent", controllers.GetBoardContent)

	router.POST("/updateProjectContent", controllers.UpdateProject)

	return router
}