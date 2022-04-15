package model

type Data struct {
	URL            string  `json:"url"`
	Views          int     `json:"views"`
	RelevanceScore float64 `json:"relevanceScore"`
}

type BodyGenerated struct {
	Data []Data `json:"data"`
}
