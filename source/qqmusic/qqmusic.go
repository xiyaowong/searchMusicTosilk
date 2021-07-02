package qqmusic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"searchMusicTosilk/source"
)

// QQMusic qq音乐
func QQMusic(songName string) (ret *source.Result, err error) {
	searchURL := "https://c.y.qq.com/soso/fcgi-bin/client_search_cp"
	fullURL := fmt.Sprintf("%s?ct=24&qqmusic_ver=1298&new_json=1&remoteplace=txt.yqq.song&searchid=&t=0&aggr=1&cr=1&catZhida=1&lossless=0&flag_qc=0&p=1&n=1&w=%s&format=json", searchURL, songName)
	req, _ := http.NewRequest("GET", fullURL, strings.NewReader(""))
	resp, err := source.Client.Do(req)
	if err != nil {
		// log.Printf("request err: %v\n", err)
		log.Println("QQMusic 请求失败")
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
	if info.Code == 0 {
		ret := &source.Result{}
		song := info.Data.Song.List[0]
		singers := []string{}
		for _, singerInfo := range song.Singer {
			singers = append(singers, singerInfo.Name)
		}
		ret.Author = strings.Join(singers, "&")
		ret.SongName = song.Name
		songURL, err := getPlayURL(song.Mid)
		if err != nil {
			return nil, err
		}
		ret.SongURL = songURL
		return ret, nil
	}
	return nil, fmt.Errorf("not found")
}

func getPlayURL(mid string) (playURL string, err error) {
	vkeyURL := "https://u.y.qq.com/cgi-bin/musicu.fcg"
	param := "{%22req%22:%20{%22module%22:%20%22CDN.SrfCdnDispatchServer%22,%20%22method%22:%20%22GetCdnDispatch%22,%20%22param%22:%20{%22guid%22:%20%223982823384%22,%20%22calltype%22:%200,%20%22userip%22:%20%22%22}},%20%22req_0%22:%20{%22module%22:%20%22vkey.GetVkeyServer%22,%20%22method%22:%20%22CgiGetVkey%22,%20%22param%22:%20{%22guid%22:%20%223982823384%22,%20%22songmid%22:%20[%22" + mid + "%22],%20%22songtype%22:%20[0],%20%22uin%22:%20%220%22,%20%22loginflag%22:%201,%20%22platform%22:%20%2220%22}},%20%22comm%22:%20{%22uin%22:%200,%20%22format%22:%20%22json%22,%20%22ct%22:%2024,%20%22cv%22:%200}}"
	fullURL := vkeyURL + "?data=" + param
	req, _ := http.NewRequest("GET", fullURL, strings.NewReader(""))

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

	info := &VkeyJSONStruct{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Printf("json unmarshal err: %v\n", err)
		return
	}
	if info.Code == 0 {
		if purl := info.Req0.Data.Midurlinfo[0].Purl; len(purl) > 0 {
			return "https://isure.stream.qqmusic.qq.com/" + purl, nil
		}
	}
	return "", fmt.Errorf("not found")
}
