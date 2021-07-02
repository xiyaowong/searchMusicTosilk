package kuwo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"searchMusicTosilk/source"
)

// Kuwo kuwo 音乐
func Kuwo(songName string) (ret *source.Result, err error) {
	u, _ := url.ParseRequestURI("http://kuwo.cn/api/www/search/searchMusicBykeyWord")
	query := url.Values{}
	query.Set("key", songName)
	query.Set("pn", "1")
	query.Set("rn", "1")
	u.RawQuery = query.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Host", "kuwo.cn")
	req.Header.Add("Referer", "http://kuwo.cn/")
	req.Header.Add("csrf", "4IT871VN3DA")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36 Edg/84.0.522.48")
	req.Header.Add("Cookie", "kw_token=4IT871VN3DA")

	resp, err := source.Client.Do(req)
	if err != nil {
		log.Printf("request err: %v\n", err)
		return nil, fmt.Errorf("酷我网络请求失败")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body err: %v\n", err)
		return nil, fmt.Errorf("酷我获取数据失败")
	}

	info := &SearchJSONStruct{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Printf("json unmarshal err: %v\n", err)
		return nil, fmt.Errorf("获取酷我歌曲信息失败")
	}
	if info.Code == 200 {
		ret := &source.Result{}
		song := info.Data.List[0]
		ret.Author = song.Artist
		ret.SongName = song.Name
		songURL, err := getPlayURL(song.Rid)
		if err != nil {
			return nil, err
		}
		ret.SongURL = songURL
		return ret, nil
	}
	return nil, fmt.Errorf("获取酷我数据失败")
}

func getPlayURL(rid int) (playURL string, err error) {
	formats := []string{"320kmp3", "192kmp3", "128kmp3"}
	// kuwo有点奇怪，api时不时就没响应
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	for _, format := range formats {
		u := fmt.Sprintf("http://www.kuwo.cn/url?format=mp3&rid=%d&response=url&type=convert_url3&br=%s&from=web", rid, format)
		req, _ := http.NewRequest("GET", u, nil)

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		info := &GetPlayURLJSONStruct{}
		err = json.Unmarshal(body, &info)
		if err != nil {
			continue
		}
		if info.Code == 200 {
			return info.URL, nil
		}
	}
	return "", fmt.Errorf("获取歌曲链接失败")
}
