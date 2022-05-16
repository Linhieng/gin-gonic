/* 不管是发送 query 还是发送 json 格式的参数，都识别为 query 参数 */

package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.Bind(&person) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}

	if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
    fmt.Printf("\n\n%#v\n\n", c)
    fmt.Printf("\n\n%#v\n\n", c.Bind(&person))
    fmt.Printf("\n\n%#v\n\n", c.BindJSON(&person))
	c.String(200, "Success")
}