package server

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/notAI-tech/verifytweet-go/internal/pkg/ocr"
	"github.com/notAI-tech/verifytweet-go/internal/pkg/search"
	"github.com/notAI-tech/verifytweet-go/internal/pkg/text"
)

// New initialises an new instance of application server
func New() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"data": "OK"})
	})

	router.MaxMultipartMemory = 4 << 20 // 4 MiB
	v1 := router.Group("/v1")
	{
		v1.POST("/verify", func(c *gin.Context) {
			file, _, err := c.Request.FormFile("tweetImage")
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(500)
				return
			}
			defer file.Close()
			buffer := bytes.NewBuffer(nil)
			if _, err := io.Copy(buffer, file); err != nil {
				log.Print(err)
				c.AbortWithStatus(500)
				return
			}
			imageBlob, err := ocr.Rescale(buffer.Bytes())
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(500)
				return
			}
			rawText, err := ocr.ConvertToText(imageBlob)
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(500)
				return
			}
			entities, err := text.Parse(rawText)
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(500)
				return
			}
			tweets, err := search.SelfHosted(entities)
			if err != nil {
				log.Print(err)
				c.AbortWithStatus(500)
				return
			}
			simMatrix := text.CalculateSimilarityMatrix(tweets, entities)
			c.JSON(http.StatusOK, map[string]interface{}{"data": simMatrix})
		})
	}

	return router
}
