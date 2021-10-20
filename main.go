package main

import (
	"bugly/bugly"
	"fmt"
)

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, Geektutu")
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080

	config := bugly.GetDingdingConfig("./config.yml")
	fmt.Println(config.Dingding.DingdingAccessToken)

}
