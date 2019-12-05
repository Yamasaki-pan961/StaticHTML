package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	fmt.Printf("サーバー動きます")
	router:= gin.Default()
	router.LoadHTMLGlob("*.html")

	router.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html", gin.H{})
	})
	router.Run(":" + os.Getenv("PORT"))
}