package main

import "github.com/gin-gonic/gin"
import "github.com/io112/tiny-url-shorter/internal/urlutil"

func main() {
	router := gin.Default()
	router.POST("/short", urlutil.ShortUrl)
	router.POST("/long", urlutil.LongUrl)
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
