package model

type CategoryAndTagOutput struct {
	Categories []Categories `json:"categories"`
	Tags       []Tags       `json:"tags"`
}
