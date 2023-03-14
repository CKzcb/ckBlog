/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/3/14 21:58
 */

package main

import "github.com/gin-gonic/gin"

//
// main
//  @Description:
//		default engine:
//		run mode:
//		router register:
//		run info: include host and port
//
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
