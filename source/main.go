package source

import (
	"net/http"
	"time"
)

var (
	// Client 用于请求
	Client = &http.Client{
		Timeout: 10 * time.Second,
	}
)

// Result 所有搜索成功后都返回这个
type Result struct {
	Author         string `json:"author"`
	SongName       string `json:"song_name"`
	SongURL        string `json:"song_url"`
	SongSilkBase64 string `json:"song_silk_base64"`
}
