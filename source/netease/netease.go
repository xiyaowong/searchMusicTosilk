package netease

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"strconv"

	"searchMusicTosilk/source"
)

// Netease 网易云
func Netease(songName string) (ret *source.Result, err error) {
	api := "http://music.163.com/api/search/pc"
	values := url.Values{}
	values.Set("s", songName)
	values.Set("offset", "0")
	values.Set("limit", "1")
	values.Set("type", "1")

	payload := strings.NewReader(values.Encode())
	req, _ := http.NewRequest("POST", api, payload)
	req.Header.Add("Host", "music.163.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36 Edg/84.0.522.48")
	req.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	resp, err := source.Client.Do(req)
	if err != nil {
		// log.Printf("request err: %v\n", err)
		log.Println("Netease 请求失败")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Printf("read body err: %v\n", err)
		log.Println("获取数据失败")
		return
	}

	info := &JSONStruct{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Println("获取json信息失败")
		return
	}

	if info.Code == 200 && len(info.Result.Songs) != 0 {
		song := info.Result.Songs[0]
		singers := []string{}
		for _, singerInfo := range song.Artists {
			singers = append(singers, singerInfo.Name)
		}
		return &source.Result{
			Author:   strings.Join(singers, "&"),
			SongName: song.Name,
			SongURL:  fmt.Sprintf("http://music.163.com/song/media/outer/url?id=%d.mp3", song.ID),
		}, nil
	}
	return nil, fmt.Errorf("not found")
}
