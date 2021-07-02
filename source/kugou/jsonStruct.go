package kugou

// SearchJSONStruct 搜索数据结构体
type SearchJSONStruct struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Data   struct {
		Aggregation []struct {
			Key   string `json:"key"`
			Count int    `json:"count"`
		} `json:"aggregation"`
		Tab  string `json:"tab"`
		Info []struct {
			OthernameOriginal string `json:"othername_original"`
			PayType320        int    `json:"pay_type_320"`
			M4Afilesize       int    `json:"m4afilesize"`
			PriceSq           int    `json:"price_sq"`
			Isoriginal        int    `json:"isoriginal"`
			Filesize          int    `json:"filesize"`
			Source            string `json:"source"`
			Bitrate           int    `json:"bitrate"`
			Topic             string `json:"topic"`
			TransParam        struct {
				Cid              int `json:"cid"`
				DisplayRate      int `json:"display_rate"`
				CpyGrade         int `json:"cpy_grade"`
				MusicpackAdvance int `json:"musicpack_advance"`
				Exclusive        int `json:"exclusive"`
				CpyLevel         int `json:"cpy_level"`
				Display          int `json:"display"`
				PayBlockTpl      int `json:"pay_block_tpl"`
			} `json:"trans_param"`
			Price            int           `json:"price"`
			Accompany        int           `json:"Accompany"`
			OldCpy           int           `json:"old_cpy"`
			SongnameOriginal string        `json:"songname_original"`
			Singername       string        `json:"singername"`
			PayType          int           `json:"pay_type"`
			Sourceid         int           `json:"sourceid"`
			TopicURL         string        `json:"topic_url"`
			FailProcess320   int           `json:"fail_process_320"`
			PkgPrice         int           `json:"pkg_price"`
			Feetype          int           `json:"feetype"`
			Filename         string        `json:"filename"`
			Price320         int           `json:"price_320"`
			Songname         string        `json:"songname"`
			Group            []interface{} `json:"group"`
			Hash             string        `json:"hash"`
			Mvhash           string        `json:"mvhash"`
			RpType           string        `json:"rp_type"`
			Privilege        int           `json:"privilege"`
			AlbumAudioID     int           `json:"album_audio_id"`
			RpPublish        int           `json:"rp_publish"`
			AlbumID          string        `json:"album_id"`
			Ownercount       int           `json:"ownercount"`
			FoldType         int           `json:"fold_type"`
			AudioID          int           `json:"audio_id"`
			PkgPriceSq       int           `json:"pkg_price_sq"`
			Three20Filesize  int           `json:"320filesize"`
			Isnew            int           `json:"isnew"`
			Duration         int           `json:"duration"`
			PkgPrice320      int           `json:"pkg_price_320"`
			Srctype          int           `json:"srctype"`
			FailProcessSq    int           `json:"fail_process_sq"`
			Sqfilesize       int           `json:"sqfilesize"`
			FailProcess      int           `json:"fail_process"`
			Three20Hash      string        `json:"320hash"`
			Extname          string        `json:"extname"`
			Sqhash           string        `json:"sqhash"`
			PayTypeSq        int           `json:"pay_type_sq"`
			Three20Privilege int           `json:"320privilege"`
			Sqprivilege      int           `json:"sqprivilege"`
			AlbumName        string        `json:"album_name"`
			Othername        string        `json:"othername"`
		} `json:"info"`
		Correctiontype  int    `json:"correctiontype"`
		Timestamp       int    `json:"timestamp"`
		Allowerr        int    `json:"allowerr"`
		Total           int    `json:"total"`
		Istag           int    `json:"istag"`
		Istagresult     int    `json:"istagresult"`
		Forcecorrection int    `json:"forcecorrection"`
		Correctiontip   string `json:"correctiontip"`
	} `json:"data"`
	Errcode int `json:"errcode"`
}

// SongInfoJSONStruct 歌曲信息
type SongInfoJSONStruct struct {
	Status  int `json:"status"`
	ErrCode int `json:"err_code"`
	Data    struct {
		Hash       string      `json:"hash"`
		Timelength int         `json:"timelength"`
		Filesize   int         `json:"filesize"`
		AudioName  string      `json:"audio_name"`
		HaveAlbum  int         `json:"have_album"`
		AlbumName  string      `json:"album_name"`
		AlbumID    string      `json:"album_id"`
		Img        string      `json:"img"`
		HaveMv     int         `json:"have_mv"`
		VideoID    interface{} `json:"video_id"`
		AuthorName string      `json:"author_name"`
		SongName   string      `json:"song_name"`
		Lyrics     string      `json:"lyrics"`
		AuthorID   string      `json:"author_id"`
		Privilege  int         `json:"privilege"`
		Privilege2 string      `json:"privilege2"`
		PlayURL    string      `json:"play_url"`
		Authors    []struct {
			AuthorID      string `json:"author_id"`
			IsPublish     string `json:"is_publish"`
			SizableAvatar string `json:"sizable_avatar"`
			AuthorName    string `json:"author_name"`
			Avatar        string `json:"avatar"`
		} `json:"authors"`
		IsFreePart       int    `json:"is_free_part"`
		Bitrate          int    `json:"bitrate"`
		RecommendAlbumID string `json:"recommend_album_id"`
		AudioID          string `json:"audio_id"`
		PlayBackupURL    string `json:"play_backup_url"`
	} `json:"data"`
}
