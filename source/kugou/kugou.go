package kugou

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"searchMusicTosilk/source"
)

// Kugou kugou音乐
func Kugou(songName string) (ret *source.Result, err error) {
	fullURL := fmt.Sprintf("http://msearchcdn.kugou.com/api/v3/search/song?pagesize=1&keyword=%s", songName)
	req, _ := http.NewRequest("GET", fullURL, nil)
	req.Header.Add("Host", "msearchcdn.kugou.com")
	resp, err := source.Client.Do(req)
	if err != nil {
		// log.Printf("request err: %v\n", err)
		log.Println("kugou 请求失败")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Printf("read body err: %v\n", err)
		log.Println("获取数据失败")
		return
	}

	info := &SearchJSONStruct{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		// log.Printf("json unmarshal err: %v\n", err)
		log.Println("获取json信息失败")
		return
	}
	if info.Status == 1 && info.Error == "" {
		ret := &source.Result{}
		song := info.Data.Info[0]
		ret.Author = song.Singername
		ret.SongName = song.Songname
		songURL, err := getPlayURL(song.Hash)
		if err != nil {
			return nil, err
		}
		ret.SongURL = songURL
		return ret, nil
	}
	return nil, fmt.Errorf("not found")
}

func getPlayURL(hash string) (playURL string, err error) {
	fullURL := fmt.Sprintf("https://wwwapi.kugou.com/yy/index.php?r=play/getdata&dfid=2kuKRO3GStCZ0VBY9V12pXeT&mid=f679eeece44cf6bec74d2867be4901f7&hash=%s", hash)

	req, _ := http.NewRequest("GET", fullURL, nil)
	req.Header.Add("Host", "wwwapi.kugou.com")

	resp, err := source.Client.Do(req)
	if err != nil {
		log.Printf("request err: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body err: %v\n", err)
		return
	}

	info := &SongInfoJSONStruct{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Printf("json unmarshal err: %v\n", err)
		return
	}
	if info.Status == 1 && info.ErrCode == 0 {
		return info.Data.PlayURL, nil
	}
	return "", fmt.Errorf("not found")
}
