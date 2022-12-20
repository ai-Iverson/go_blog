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

type SiteOutput struct {
	Badges         []Badges         `json:"badges"`
	SiteInfo       SiteInfo         `json:"siteInfo"`
	Introduction   Introduction     `json:"introduction"`
	NewBlogList    []NewBlogList    `json:"newBlogList"`
	CategoryList   []Categories     `json:"categoryList"`
	TagList        []Tags           `json:"tagList"`
	RandomBlogList []RandomBlogList `json:"randomBlogList"`
}

type Badges struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Subject string `json:"subject"`
	Value   string `json:"value"`
	Color   string `json:"color"`
}

type NewBlogList struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Password string `json:"password"`
	Privacy  bool   `json:"privacy"`
}

type RandomBlogList struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	FirstPicture string `json:"firstPicture"`
	CreateTime   string `json:"createTime"`
	Password     string `json:"password"`
	Privacy      bool   `json:"privacy"`
}

type Introduction struct {
	Avatar    string      `json:"avatar"`
	Name      string      `json:"name"`
	Github    string      `json:"github"`
	Telegram  string      `json:"telegram"`
	Qq        string      `json:"qq"`
	Bilibili  string      `json:"bilibili"`
	Netease   string      `json:"netease"`
	Email     string      `json:"email"`
	RollText  []string    `json:"rollText"`
	Favorites []Favorites `json:"favorites"`
}

type Favorites struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SiteInfo struct {
	Reward           string    `json:"reward"`
	Copyright        Copyright `json:"copyright"`
	BlogName         string    `json:"blogName"`
	Beian            string    `json:"beian"`
	WebTitleSuffix   string    `json:"webTitleSuffix"`
	FooterImgTitle   string    `json:"footerImgTitle"`
	CommentAdminFlag string    `json:"commentAdminFlag"`
	FooterImgURL     string    `json:"footerImgUrl"`
}

type Copyright struct {
	Title    string `json:"title"`
	SiteName string `json:"siteName"`
}
