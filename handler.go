package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"searchMusicTosilk/silk"
	"searchMusicTosilk/source"
	"searchMusicTosilk/source/kugou"
	"searchMusicTosilk/source/migu"
	"searchMusicTosilk/source/netease"
	"searchMusicTosilk/source/qqmusic"
	"searchMusicTosilk/source/kuwo"
)

// Response 响应模板
type Response struct {
	Code   int            `json:"code"`
	Msg    string         `json:"msg"`
	Result *source.Result `json:"result"`
}

// MusicHandler 歌曲
func MusicHandler(c *gin.Context) {
	option := c.Query("option")
	name := c.Query("name")
	if len(strings.TrimSpace(name)) == 0 {
		c.JSON(http.StatusOK, &Response{
			Code: http.StatusBadRequest,
			Msg:  "缺少参数[name]",
		})
		return
	}

	var sourcer func(string) (*source.Result, error)
	switch option {
	case "qq":
		sourcer = qqmusic.QQMusic
	case "netease":
		sourcer = netease.Netease
	case "migu":
		sourcer = migu.Migu
	case "kugou":
		sourcer = kugou.Kugou
	case "kuwo":
		sourcer = kuwo.Kuwo
	default:
		c.JSON(http.StatusOK, &Response{
			Code: 1,
			Msg:  "无效的参数[option]",
		})
		return
	}
	ret, err := sourcer(name)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, &Response{
			Code: 1,
			Msg:  "出错",
		})
		return
	}
	// 转silk?
	if c.Query("silk") != "" {
		if b64, err := silk.URL2SilkBase64(ret.SongURL); err == nil {
			ret.SongSilkBase64 = b64
		}
	}
	c.JSON(http.StatusOK, &Response{
		Code:   0,
		Msg:    "",
		Result: ret,
	})
}

// SilkHandler silk编码
func SilkHandler(c *gin.Context) {
	fileURL := c.Query("url")
	if !govalidator.IsURL(fileURL) {
		c.JSON(http.StatusOK,
			&Response{
				Code: 1,
				Msg:  "请输入有效链接",
			})
		return
	}
	fmt.Printf("取得音频链接: [%s]\n", fileURL)
	b64, err := silk.URL2SilkBase64(fileURL)
	if err != nil {
		log.Println("url转base64失败")
		c.JSON(http.StatusOK,
			&Response{
				Code: 1,
				Msg:  fmt.Sprintf("url转base64失败, 请确保传入url正确[%s]", fileURL),
			})
		return
	}
	c.JSON(http.StatusOK,
		&Response{
			Code: 0,
			Result: &source.Result{
				SongSilkBase64: b64,
			}})
}
