package models

type CiklumResponse struct {
	HttpStatus int       `json:"httpStatus"`
	Response   CResponse `json:"response"`
}

type CResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Type              string  `json:"type"`
	HarvesterId       string  `json:"harvesterId,omitempty"`
	CerebroScore      float64 `json:"cerebro-score,omitempty"`
	Url               string  `json:"url,omitempty"`
	Title             string  `json:"title,omitempty"`
	CleanImage        string  `json:"cleanImage,omitempty"`
	CommercialPartner string  `json:"commercialPartner,omitempty"`
	LogoURL           string  `json:"logoURL,omitempty"`
}
