package migu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"searchMusicTosilk/source"
)

// Migu 咪咕音乐
func Migu(songName string) (ret *source.Result, err error) {
	u, _ := url.ParseRequestURI("http://pd.musicapp.migu.cn/MIGUM3.0/v1.0/content/search_all.do")
	query := url.Values{}
	query.Set("pageNo", "1")
	query.Set("pageSize", "1")
	query.Set("searchSwitch", `{"song":1,"album":0,"singer":0,"tagSong":0,"mvSong":0,"songlist":0,"bestShow":0}`)
	query.Set("text", songName)
	u.RawQuery = query.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)

	req.Header.Add("Referer", "https://m.music.migu.cn/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Mobile Safari/537.36")
	req.Header.Add("Host", "https://migu.cn")

	resp, err := source.Client.Do(req)
	if err != nil {
		// log.Printf("request err: %v\n", err)
		log.Println("Migu 请求失败，多半是ip被ban了.")
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

	if info.Code == "000000" && len(info.SongResultData.Result) > 0 {
		song := info.SongResultData.Result[0]
		singers := []string{}
		for _, singerInfo := range song.Singers {
			singers = append(singers, singerInfo.Name)
		}
		return &source.Result{
			Author:   strings.Join(singers, "&"),
			SongName: song.Name,
			SongURL:  fmt.Sprintf("https://app.pd.nf.migu.cn/MIGUM3.0/v1.0/content/sub/listenSong.do?channel=mx&copyrightId=%s&contentId=%s&toneFlag=LQ&resourceType=2&userId=15548614588710179085069&netType=00", song.CopyrightID, song.ContentID),
		}, nil
	}
	return nil, fmt.Errorf("not found")
}
