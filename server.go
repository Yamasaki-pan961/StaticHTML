package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main(){
	fmt.Printf("サーバー動きます")
	router:= gin.Default()
	router.LoadHTMLGlob("*.html")

	router.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html", gin.H{})
	})	
	fmt.Printf("サーバー動きました")
	router.Run(":" + os.Getenv("PORT"))

}