package main

import (
	"bugly/bugly"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config := bugly.GetDingdingConfig("./config.yml")
	fmt.Println(config.Dingding.DingdingAccessToken)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.POST("/", func(c *gin.Context) {
		json := map[string]interface{}{}
		c.BindJSON(&json)
		markdownString := bugly.ParseJsonString(json)
		fmt.Println(markdownString)
		bugly.SendDingDingMessage("每日Crash统计", markdownString, config.Dingding.DingdingURL, config.Dingding.DingdingAccessToken)
		c.String(200, "ok")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
