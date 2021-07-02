package migu

// JSONStruct 咪咕返回json
type JSONStruct struct {
	Code           string `json:"code"`
	Info           string `json:"info"`
	SongResultData struct {
		TotalCount  string        `json:"totalCount"`
		Correct     []interface{} `json:"correct"`
		ResultType  string        `json:"resultType"`
		IsFromCache string        `json:"isFromCache"`
		Result      []struct {
			ID           string   `json:"id"`
			ResourceType string   `json:"resourceType"`
			ContentID    string   `json:"contentId"`
			CopyrightID  string   `json:"copyrightId"`
			Name         string   `json:"name"`
			HighlightStr []string `json:"highlightStr"`
			Singers      []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"singers"`
			Albums []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"albums"`
			Tags     []string `json:"tags"`
			LyricURL string   `json:"lyricUrl"`
			TrcURL   string   `json:"trcUrl"`
			ImgItems []struct {
				ImgSizeType string `json:"imgSizeType"`
				Img         string `json:"img"`
			} `json:"imgItems"`
			TelevisionNames []string `json:"televisionNames"`
			Tones           []struct {
				ID          string `json:"id"`
				CopyrightID string `json:"copyrightId"`
				Price       string `json:"price"`
				ExpireDate  string `json:"expireDate"`
			} `json:"tones"`
			RelatedSongs []struct {
				ResourceType     string `json:"resourceType"`
				ResourceTypeName string `json:"resourceTypeName"`
				CopyrightID      string `json:"copyrightId"`
				ProductID        string `json:"productId"`
			} `json:"relatedSongs"`
			ToneControl string `json:"toneControl"`
			RateFormats []struct {
				ResourceType         string `json:"resourceType"`
				FormatType           string `json:"formatType"`
				URL                  string `json:"url,omitempty"`
				Format               string `json:"format"`
				Size                 string `json:"size"`
				FileType             string `json:"fileType,omitempty"`
				Price                string `json:"price"`
				IosURL               string `json:"iosUrl,omitempty"`
				AndroidURL           string `json:"androidUrl,omitempty"`
				IosSize              string `json:"iosSize,omitempty"`
				AndroidSize          string `json:"androidSize,omitempty"`
				IosFormat            string `json:"iosFormat,omitempty"`
				AndroidFormat        string `json:"androidFormat,omitempty"`
				IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
				AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
			} `json:"rateFormats"`
			NewRateFormats []struct {
				ResourceType         string `json:"resourceType"`
				FormatType           string `json:"formatType"`
				URL                  string `json:"url,omitempty"`
				Format               string `json:"format"`
				Size                 string `json:"size"`
				FileType             string `json:"fileType,omitempty"`
				Price                string `json:"price"`
				IosURL               string `json:"iosUrl,omitempty"`
				AndroidURL           string `json:"androidUrl,omitempty"`
				IosSize              string `json:"iosSize,omitempty"`
				AndroidSize          string `json:"androidSize,omitempty"`
				IosFormat            string `json:"iosFormat,omitempty"`
				AndroidFormat        string `json:"androidFormat,omitempty"`
				IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
				AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
			} `json:"newRateFormats"`
			SongType         string `json:"songType"`
			IsInDAlbum       string `json:"isInDAlbum"`
			Copyright        string `json:"copyright"`
			DigitalColumnID  string `json:"digitalColumnId"`
			Mrcurl           string `json:"mrcurl"`
			SongDescs        string `json:"songDescs"`
			SongAliasName    string `json:"songAliasName"`
			InvalidateDate   string `json:"invalidateDate"`
			IsInSalesPeriod  string `json:"isInSalesPeriod"`
			DalbumID         string `json:"dalbumId"`
			IsInSideDalbum   string `json:"isInSideDalbum"`
			VipType          string `json:"vipType"`
			ChargeAuditions  string `json:"chargeAuditions"`
			ScopeOfcopyright string `json:"scopeOfcopyright"`
			ActionURLParams  string `json:"actionUrlParams"`
			MgVideoParam4IOS string `json:"mgVideoParam4IOS"`
			MgVideoParam4Adr string `json:"mgVideoParam4Adr"`
		} `json:"result"`
		TipStatus string `json:"tipStatus"`
	} `json:"songResultData"`
	BestShowResultToneData struct {
	} `json:"bestShowResultToneData"`
}
