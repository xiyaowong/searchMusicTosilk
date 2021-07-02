package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"searchMusicTosilk/asset"
)

func init() {
}

func main() {
	port := flag.Int("port", 3317, "运行的端口号")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	// router := gin.Default()
	router := gin.New()

	// 搜歌
	router.GET("/music", MusicHandler)
	// 转silk编码
	router.GET("/silk", SilkHandler)

	// 网站图标
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Writer.Header().Add("Content-Encoding", "gzip")
		c.Writer.Header().Add("Content-Type", "image/png")
		c.Writer.Write(asset.FaviconIco)
	})

	// 其他一律是使用说明页面
	router.NoRoute(func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "text/html")
		c.Writer.Write(asset.Doc)
	})

	fmt.Printf("Running on 0.0.0.0%s\n", fmt.Sprintf(":%d", *port))
	router.Run(fmt.Sprintf(":%d", *port))
}
