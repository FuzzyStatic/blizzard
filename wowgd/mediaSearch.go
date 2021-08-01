package wowgd

// MediaSearch Structure
type MediaSearch struct {
	Page              int  `json:"page"`
	PageSize          int  `json:"pageSize"`
	MaxPageSize       int  `json:"maxPageSize"`
	PageCount         int  `json:"pageCount"`
	ResultCountCapped bool `json:"resultCountCapped"`
	Results           []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
			Assets []struct {
				FileDataID int    `json:"file_data_id"`
				Value      string `json:"value"`
				Key        string `json:"key"`
			} `json:"assets"`
			ID int `json:"id"`
		} `json:"data"`
	} `json:"results"`
}
