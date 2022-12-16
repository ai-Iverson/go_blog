package model

type SiteSettingsDetilOutput struct {
	Type3 []Type3 `json:"type3"`
	Type2 []Type2 `json:"type2"`
	Type1 []Type1 `json:"type1"`
}
type Type3 struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}
type Type2 struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}
type Type1 struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}

type UpdateSiteSettingsInput struct {
	Settings  []Settings `json:"settings"`
	DeleteIds []int      `json:"deleteIds"`
}

type Settings struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}
