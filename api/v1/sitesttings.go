package v1

import "github.com/gogf/gf/v2/frame/g"

type SiteSettingsDetilReq struct {
	g.Meta        `path:"/admin/siteSettings" method:"get" summart:" 获取站点信息" tags:"获取站点信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type SiteSettingsDetilRes struct {
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

type UpdateSiteSettingsReq struct {
	g.Meta        `path:"/admin/siteSettings" method:"post" summart:" 获取站点信息" tags:"获取站点信息"`
	Authorization string     `json:"Authorization" in:"header"  dc:"Authorization"`
	Settings      []Settings `json:"settings"`
	DeleteIds     []int      `json:"deleteIds"`
}

type Settings struct {
	ID     int    `json:"id"`
	NameEn string `json:"nameEn"`
	NameZh string `json:"nameZh"`
	Value  string `json:"value"`
	Type   int    `json:"type"`
}

type UpdateSiteSettingsRes struct {
}

type SiteReq struct {
	g.Meta `path:"/site" method:"get" summart:" 获取站点信息" tags:"获取站点信息"`
}

type SiteRes struct {
	Badges         []Badges         `json:"badges"`
	SiteInfo       SiteInfo         `json:"siteInfo"`
	Introduction   Introduction     `json:"introduction"`
	NewBlogList    []NewBlogList    `json:"newBlogList"`
	CategoryList   []Categorys      `json:"categoryList"`
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
