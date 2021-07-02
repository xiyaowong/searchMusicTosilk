package netease

// JSONStruct 网易信息
type JSONStruct struct {
	Result Result `json:"result"`
	Code   int    `json:"code"`
}
type Artists struct {
	Name      string        `json:"name"`
	ID        int           `json:"id"`
	PicID     int           `json:"picId"`
	Img1V1ID  int           `json:"img1v1Id"`
	BriefDesc string        `json:"briefDesc"`
	PicURL    string        `json:"picUrl"`
	Img1V1URL string        `json:"img1v1Url"`
	AlbumSize int           `json:"albumSize"`
	Alias     []interface{} `json:"alias"`
	Trans     string        `json:"trans"`
	MusicSize int           `json:"musicSize"`
}
type Artist struct {
	Name      string        `json:"name"`
	ID        int           `json:"id"`
	PicID     int           `json:"picId"`
	Img1V1ID  int           `json:"img1v1Id"`
	BriefDesc string        `json:"briefDesc"`
	PicURL    string        `json:"picUrl"`
	Img1V1URL string        `json:"img1v1Url"`
	AlbumSize int           `json:"albumSize"`
	Alias     []interface{} `json:"alias"`
	Trans     string        `json:"trans"`
	MusicSize int           `json:"musicSize"`
}
type Album struct {
	Name            string        `json:"name"`
	ID              int           `json:"id"`
	Type            string        `json:"type"`
	Size            int           `json:"size"`
	PicID           int64         `json:"picId"`
	BlurPicURL      string        `json:"blurPicUrl"`
	CompanyID       int           `json:"companyId"`
	Pic             int64         `json:"pic"`
	PicURL          string        `json:"picUrl"`
	PublishTime     int64         `json:"publishTime"`
	Description     string        `json:"description"`
	Tags            string        `json:"tags"`
	Company         string        `json:"company"`
	BriefDesc       string        `json:"briefDesc"`
	Artist          Artist        `json:"artist"`
	Songs           []interface{} `json:"songs"`
	Alias           []interface{} `json:"alias"`
	Status          int           `json:"status"`
	CopyrightID     int           `json:"copyrightId"`
	CommentThreadID string        `json:"commentThreadId"`
	Artists         []Artists     `json:"artists"`
	PicIDStr        string        `json:"picId_str"`
}
type HMusic struct {
	Name        interface{} `json:"name"`
	ID          int         `json:"id"`
	Size        int         `json:"size"`
	Extension   string      `json:"extension"`
	Sr          int         `json:"sr"`
	DfsID       int         `json:"dfsId"`
	Bitrate     int         `json:"bitrate"`
	PlayTime    int         `json:"playTime"`
	VolumeDelta float64     `json:"volumeDelta"`
}
type MMusic struct {
	Name        interface{} `json:"name"`
	ID          int         `json:"id"`
	Size        int         `json:"size"`
	Extension   string      `json:"extension"`
	Sr          int         `json:"sr"`
	DfsID       int         `json:"dfsId"`
	Bitrate     int         `json:"bitrate"`
	PlayTime    int         `json:"playTime"`
	VolumeDelta float64     `json:"volumeDelta"`
}
type LMusic struct {
	Name        interface{} `json:"name"`
	ID          int         `json:"id"`
	Size        int         `json:"size"`
	Extension   string      `json:"extension"`
	Sr          int         `json:"sr"`
	DfsID       int         `json:"dfsId"`
	Bitrate     int         `json:"bitrate"`
	PlayTime    int         `json:"playTime"`
	VolumeDelta float64     `json:"volumeDelta"`
}
type BMusic struct {
	Name        interface{} `json:"name"`
	ID          int         `json:"id"`
	Size        int         `json:"size"`
	Extension   string      `json:"extension"`
	Sr          int         `json:"sr"`
	DfsID       int         `json:"dfsId"`
	Bitrate     int         `json:"bitrate"`
	PlayTime    int         `json:"playTime"`
	VolumeDelta float64     `json:"volumeDelta"`
}
type Songs struct {
	Name            string        `json:"name"`
	ID              int           `json:"id"`
	Position        int           `json:"position"`
	Alias           []interface{} `json:"alias"`
	Status          int           `json:"status"`
	Fee             int           `json:"fee"`
	CopyrightID     int           `json:"copyrightId"`
	Disc            string        `json:"disc"`
	No              int           `json:"no"`
	Artists         []Artists     `json:"artists"`
	Album           Album         `json:"album"`
	Starred         bool          `json:"starred"`
	Popularity      float64       `json:"popularity"`
	Score           int           `json:"score"`
	StarredNum      int           `json:"starredNum"`
	Duration        int           `json:"duration"`
	PlayedNum       int           `json:"playedNum"`
	DayPlays        int           `json:"dayPlays"`
	HearTime        int           `json:"hearTime"`
	Ringtone        interface{}   `json:"ringtone"`
	Crbt            interface{}   `json:"crbt"`
	Audition        interface{}   `json:"audition"`
	CopyFrom        string        `json:"copyFrom"`
	CommentThreadID string        `json:"commentThreadId"`
	RtURL           interface{}   `json:"rtUrl"`
	Ftype           int           `json:"ftype"`
	RtUrls          []interface{} `json:"rtUrls"`
	Copyright       int           `json:"copyright"`
	HMusic          HMusic        `json:"hMusic"`
	MMusic          MMusic        `json:"mMusic"`
	LMusic          LMusic        `json:"lMusic"`
	BMusic          BMusic        `json:"bMusic"`
	Mp3URL          string        `json:"mp3Url"`
	Mvid            int           `json:"mvid"`
	Rtype           int           `json:"rtype"`
	Rurl            interface{}   `json:"rurl"`
}
type Result struct {
	Songs     []Songs `json:"songs"`
	SongCount int     `json:"songCount"`
}
