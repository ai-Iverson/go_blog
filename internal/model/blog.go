package model

import v1 "go_blog/api/v1"

type CreateBlogInput struct {
	Title            string // 文章标题
	FirstPicture     string // 文章首图，用于随机文章展示
	Content          string // 文章正文
	Description      string // 描述
	IsPublished      bool   // 公开或私密
	IsRecommend      bool   // 推荐开关
	IsAppreciation   bool   // 赞赏开关
	IsCommentEnabled bool   // 评论开关
	Views            int    // 浏览次数
	Words            string // 文章字数
	ReadTime         int    // 阅读时长(分钟)
	CategoryId       int    // 文章分类
	IsTop            bool   // 是否置顶
	Password         string // 密码保护
	User             string // 文章作者
	TagList          []int  // 标签列表
}
type Blog struct {
	ID               int         `json:"id"`
	Title            string      `json:"title"`
	FirstPicture     string      `json:"firstPicture"`
	Content          interface{} `json:"content"`
	Description      interface{} `json:"description"`
	IsPublished      bool        `json:"published"`
	IsRecommend      bool        `json:"recommend"`
	IsAppreciation   bool        `json:"appreciation"`
	IsCommentEnabled bool        `json:"commentEnabled"`
	IsTop            bool        `json:"top"`
	CreateTime       string      `json:"createTime"`
	UpdateTime       string      `json:"updateTime"`
	Views            interface{} `json:"views"`
	Words            interface{} `json:"words"`
	ReadTime         interface{} `json:"readTime"`
	Password         string      `json:"password"`
	User             interface{} `json:"user"`
	CategoryId       int64       `json:"categoryId"`
	Category         struct {
		ID           int           `json:"id"`
		CategoryName string        `json:"name"`
		Blogs        []interface{} `json:"blogs"`
	} `json:"category"`
	Tags []interface{} `json:"tags"`
}

type BlogsListOutput struct {
	List []Blog `json:"list"`
	v1.CommonPaginationReq
}

type BlogDetailOutput struct {
	ID             int         `json:"id"`
	Title          string      `json:"title"`
	FirstPicture   string      `json:"firstPicture"`
	Content        string      `json:"content"`
	Description    string      `json:"description"`
	Published      bool        `json:"published"`
	Recommend      bool        `json:"recommend"`
	Appreciation   bool        `json:"appreciation"`
	CommentEnabled bool        `json:"commentEnabled"`
	Top            bool        `json:"top"`
	CreateTime     string      `json:"createTime"`
	UpdateTime     string      `json:"updateTime"`
	Views          int         `json:"views"`
	Words          int         `json:"words"`
	ReadTime       int         `json:"readTime"`
	Password       string      `json:"password"`
	User           interface{} `json:"user"`
	Category       Categories  `json:"category"`
	Tags           []Tags      `json:"tags"`
}
