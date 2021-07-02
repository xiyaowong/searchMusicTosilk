package kuwo

// SearchJSONStruct 搜索
type SearchJSONStruct struct {
	Code    int   `json:"code"`
	CurTime int64 `json:"curTime"`
	Data    struct {
		Total string `json:"total"`
		List  []struct {
			Musicrid  string `json:"musicrid"`
			Artist    string `json:"artist"`
			Mvpayinfo struct {
				Play int `json:"play"`
				Vid  int `json:"vid"`
				Down int `json:"down"`
			} `json:"mvpayinfo"`
			Pic             string `json:"pic"`
			Isstar          int    `json:"isstar"`
			Rid             int    `json:"rid"`
			Duration        int    `json:"duration"`
			Score100        string `json:"score100"`
			ContentType     string `json:"content_type"`
			Track           int    `json:"track"`
			HasLossless     bool   `json:"hasLossless"`
			Hasmv           int    `json:"hasmv"`
			ReleaseDate     string `json:"releaseDate"`
			Album           string `json:"album"`
			Albumid         int    `json:"albumid"`
			Pay             string `json:"pay"`
			Artistid        int    `json:"artistid"`
			Albumpic        string `json:"albumpic"`
			SongTimeMinutes string `json:"songTimeMinutes"`
			IsListenFee     bool   `json:"isListenFee"`
			Pic120          string `json:"pic120"`
			Name            string `json:"name"`
			Online          int    `json:"online"`
			PayInfo         struct {
				Play             string `json:"play"`
				Download         string `json:"download"`
				LocalEncrypt     string `json:"local_encrypt"`
				CannotDownload   int    `json:"cannotDownload"`
				CannotOnlinePlay int    `json:"cannotOnlinePlay"`
				FeeType          struct {
					Song string `json:"song"`
					Vip  string `json:"vip"`
				} `json:"feeType"`
				Down string `json:"down"`
			} `json:"payInfo"`
		} `json:"list"`
	} `json:"data"`
	Msg       string `json:"msg"`
	ProfileID string `json:"profileId"`
	ReqID     string `json:"reqId"`
}

// GetPlayURLJSONStruct 获取播放链接
type GetPlayURLJSONStruct struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	URL  string `json:"url"`
}