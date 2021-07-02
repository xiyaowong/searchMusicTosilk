// Package jsonStruct provides ...
package qqmusic

// VkeyJSONStruct qq音乐获取链接信息
type VkeyJSONStruct struct {
	Code    int   `json:"code"`
	Ts      int64 `json:"ts"`
	StartTs int64 `json:"start_ts"`
	Req     Req   `json:"req"`
	Req0    Req0  `json:"req_0"`
}
type DataReq struct {
	Expiration    int      `json:"expiration"`
	Freeflowsip   []string `json:"freeflowsip"`
	Keepalivefile string   `json:"keepalivefile"`
	Msg           string   `json:"msg"`
	Retcode       int      `json:"retcode"`
	Servercheck   string   `json:"servercheck"`
	Sip           []string `json:"sip"`
	Testfile2G    string   `json:"testfile2g"`
	Testfilewifi  string   `json:"testfilewifi"`
	Uin           string   `json:"uin"`
	Userip        string   `json:"userip"`
	Vkey          string   `json:"vkey"`
}
type Req struct {
	Code int     `json:"code"`
	Data DataReq `json:"data"`
}
type Midurlinfo struct {
	CommonDownfromtag int    `json:"common_downfromtag"`
	Errtype           string `json:"errtype"`
	Filename          string `json:"filename"`
	Flowfromtag       string `json:"flowfromtag"`
	Flowurl           string `json:"flowurl"`
	Hisbuy            int    `json:"hisbuy"`
	Hisdown           int    `json:"hisdown"`
	Isbuy             int    `json:"isbuy"`
	Isonly            int    `json:"isonly"`
	Onecan            int    `json:"onecan"`
	Opi128Kurl        string `json:"opi128kurl"`
	Opi192Koggurl     string `json:"opi192koggurl"`
	Opi192Kurl        string `json:"opi192kurl"`
	Opi30Surl         string `json:"opi30surl"`
	Opi48Kurl         string `json:"opi48kurl"`
	Opi96Kurl         string `json:"opi96kurl"`
	Opiflackurl       string `json:"opiflackurl"`
	P2Pfromtag        int    `json:"p2pfromtag"`
	Pdl               int    `json:"pdl"`
	Pneed             int    `json:"pneed"`
	Pneedbuy          int    `json:"pneedbuy"`
	Premain           int    `json:"premain"`
	Purl              string `json:"purl"`
	Qmdlfromtag       int    `json:"qmdlfromtag"`
	Result            int    `json:"result"`
	Songmid           string `json:"songmid"`
	Tips              string `json:"tips"`
	UIAlert           int    `json:"uiAlert"`
	VipDownfromtag    int    `json:"vip_downfromtag"`
	Vkey              string `json:"vkey"`
	Wififromtag       string `json:"wififromtag"`
	Wifiurl           string `json:"wifiurl"`
}
type DataReq0 struct {
	Expiration   int          `json:"expiration"`
	LoginKey     string       `json:"login_key"`
	Midurlinfo   []Midurlinfo `json:"midurlinfo"`
	Msg          string       `json:"msg"`
	Retcode      int          `json:"retcode"`
	Servercheck  string       `json:"servercheck"`
	Sip          []string     `json:"sip"`
	Testfile2G   string       `json:"testfile2g"`
	Testfilewifi string       `json:"testfilewifi"`
	Thirdip      []string     `json:"thirdip"`
	Uin          string       `json:"uin"`
	VerifyType   int          `json:"verify_type"`
}
type Req0 struct {
	Code int      `json:"code"`
	Data DataReq0 `json:"data"`
}

// SearchJSONStruct qq音乐搜索接口信息
type SearchJSONStruct struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Notice  string `json:"notice"`
	Subcode int    `json:"subcode"`
	Time    int    `json:"time"`
	Tips    string `json:"tips"`
}
type Semantic struct {
	Curnum   int           `json:"curnum"`
	Curpage  int           `json:"curpage"`
	List     []interface{} `json:"list"`
	Totalnum int           `json:"totalnum"`
}
type Action struct {
	Alert  int `json:"alert"`
	Icons  int `json:"icons"`
	Msg    int `json:"msg"`
	Switch int `json:"switch"`
}
type SearchAlbum struct {
	ID           int    `json:"id"`
	Mid          string `json:"mid"`
	Name         string `json:"name"`
	Pmid         string `json:"pmid"`
	Subtitle     string `json:"subtitle"`
	Title        string `json:"title"`
	TitleHilight string `json:"title_hilight"`
}
type File struct {
	B30S        int    `json:"b_30s"`
	E30S        int    `json:"e_30s"`
	MediaMid    string `json:"media_mid"`
	Size128     int    `json:"size_128"`
	Size128Mp3  int    `json:"size_128mp3"`
	Size320     int    `json:"size_320"`
	Size320Mp3  int    `json:"size_320mp3"`
	SizeAac     int    `json:"size_aac"`
	SizeApe     int    `json:"size_ape"`
	SizeDts     int    `json:"size_dts"`
	SizeFlac    int    `json:"size_flac"`
	SizeOgg     int    `json:"size_ogg"`
	SizeTry     int    `json:"size_try"`
	StrMediaMid string `json:"strMediaMid"`
	TryBegin    int    `json:"try_begin"`
	TryEnd      int    `json:"try_end"`
}
type Ksong struct {
	ID  int    `json:"id"`
	Mid string `json:"mid"`
}
type Mv struct {
	ID  int    `json:"id"`
	Vid string `json:"vid"`
}
type Pay struct {
	PayDown    int `json:"pay_down"`
	PayMonth   int `json:"pay_month"`
	PayPlay    int `json:"pay_play"`
	PayStatus  int `json:"pay_status"`
	PriceAlbum int `json:"price_album"`
	PriceTrack int `json:"price_track"`
	TimeFree   int `json:"time_free"`
}
type Singer struct {
	ID           int    `json:"id"`
	Mid          string `json:"mid"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	TitleHilight string `json:"title_hilight"`
	Type         int    `json:"type"`
	Uin          int    `json:"uin"`
}
type Volume struct {
	Gain float64 `json:"gain"`
	Lra  float64 `json:"lra"`
	Peak float64 `json:"peak"`
}
type List struct {
	Action        Action        `json:"action"`
	Album         SearchAlbum   `json:"album"`
	Chinesesinger int           `json:"chinesesinger"`
	Desc          string        `json:"desc"`
	DescHilight   string        `json:"desc_hilight"`
	Docid         string        `json:"docid"`
	Es            string        `json:"es"`
	File          File          `json:"file"`
	Fnote         int           `json:"fnote"`
	Genre         int           `json:"genre"`
	Grp           []interface{} `json:"grp"`
	ID            int           `json:"id"`
	IndexAlbum    int           `json:"index_album"`
	IndexCd       int           `json:"index_cd"`
	Interval      int           `json:"interval"`
	Isonly        int           `json:"isonly"`
	Ksong         Ksong         `json:"ksong"`
	Language      int           `json:"language"`
	Lyric         string        `json:"lyric"`
	LyricHilight  string        `json:"lyric_hilight"`
	Mid           string        `json:"mid"`
	Mv            Mv            `json:"mv"`
	Name          string        `json:"name"`
	NewStatus     int           `json:"newStatus"`
	Nt            int64         `json:"nt"`
	Ov            int           `json:"ov"`
	Pay           Pay           `json:"pay"`
	Pure          int           `json:"pure"`
	Sa            int           `json:"sa"`
	Singer        []Singer      `json:"singer"`
	Status        int           `json:"status"`
	Subtitle      string        `json:"subtitle"`
	T             int           `json:"t"`
	Tag           int           `json:"tag"`
	Tid           int           `json:"tid"`
	TimePublic    string        `json:"time_public"`
	Title         string        `json:"title"`
	TitleHilight  string        `json:"title_hilight"`
	Type          int           `json:"type"`
	URL           string        `json:"url"`
	Ver           int           `json:"ver"`
	Volume        Volume        `json:"volume"`
}
type Song struct {
	Curnum   int    `json:"curnum"`
	Curpage  int    `json:"curpage"`
	List     []List `json:"list"`
	Totalnum int    `json:"totalnum"`
}
type ZhidaMv struct {
	Desc        string `json:"desc"`
	ID          int    `json:"id"`
	Pic         string `json:"pic"`
	PublishDate string `json:"publish_date"`
	Title       string `json:"title"`
	Vid         string `json:"vid"`
}
type Zhida struct {
	Type    int     `json:"type"`
	ZhidaMv ZhidaMv `json:"zhida_mv"`
}
type Data struct {
	Keyword   string        `json:"keyword"`
	Priority  int           `json:"priority"`
	Qc        []interface{} `json:"qc"`
	Semantic  Semantic      `json:"semantic"`
	Song      Song          `json:"song"`
	Tab       int           `json:"tab"`
	Taglist   []interface{} `json:"taglist"`
	Totaltime int           `json:"totaltime"`
	Zhida     Zhida         `json:"zhida"`
}
