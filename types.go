package main

type Structure struct {
	OwnerShipData struct {
		HasMoreItems  bool `json:"hasMoreItems"`
		NumberOfItems int  `json:"numberOfItems"`
		Success       bool `json:"success"`
		Items         []struct {
			ReadStatus    string `json:"readStatus"`
			TargetDevices struct {
			} `json:"targetDevices"`
			Origin                   string        `json:"origin"`
			Title                    string        `json:"title"`
			SortableAuthors          string        `json:"sortableAuthors"`
			IsSizeGreaterThan50Mb    bool          `json:"isSizeGreaterThan50Mb"`
			Capability               []string      `json:"capability"`
			AcquiredDate             string        `json:"acquiredDate"`
			ProductImage             string        `json:"productImage"`
			CollectionList           []interface{} `json:"collectionList"`
			ContentCategoryType      string        `json:"contentCategoryType"`
			ContentType              string        `json:"contentType"`
			IsContentValid           bool          `json:"isContentValid"`
			Author                   string        `json:"author"`
			StatusFromPlatformSearch string        `json:"statusFromPlatformSearch"`
			RenderDownloadElements   bool          `json:"renderDownloadElements"`
			AcquiredTime             int64         `json:"acquiredTime"`
			SortableTitle            string        `json:"sortableTitle"`
			IsNotYetLaunched         bool          `json:"isNotYetLaunched"`
			NumericSize              int           `json:"numericSize"`
			CapabilityList           []string      `json:"capabilityList"`
			Size                     string        `json:"size"`
			CollectionCount          int           `json:"collectionCount"`
			Asin                     string        `json:"asin"`
			ContentIdentifier        string        `json:"contentIdentifier"`
			Category                 string        `json:"category"`
			NumericFileSize          int           `json:"numericFileSize"`
		} `json:"items"`
	} `json:"OwnershipData"`
}
