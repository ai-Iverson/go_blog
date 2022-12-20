package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type BlogBase struct {
	ID             int           `json:"id"`
	Title          string        `json:"title"`
	FirstPicture   string        `json:"firstPicture"`
	Content        interface{}   `json:"content"`
	Description    interface{}   `json:"description"`
	Published      bool          `json:"published"`
	Recommend      bool          `json:"recommend"`
	Appreciation   bool          `json:"appreciation"`
	CommentEnabled bool          `json:"commentEnabled"`
	Top            bool          `json:"top"`
	CreateTime     string        `json:"createTime"`
	UpdateTime     string        `json:"updateTime"`
	Views          interface{}   `json:"views"`
	Words          interface{}   `json:"words"`
	ReadTime       interface{}   `json:"readTime"`
	Password       string        `json:"password"`
	UserId         interface{}   `json:"user"`
	Category       Categorys     `json:"category"`
	Tags           []interface{} `json:"tags"`
}

type Blogs struct {
	List []BlogBase `json:"list"`
	CommonPaginationReq
}

type CreateBlogReq struct {
	g.Meta         `path:"/admin/blog" method:"post" summart:"写文章" tags:"写文章"`
	Authorization  string `json:"Authorization" in:"header"  dc:"Authorization"`
	Title          string `json:"title"`
	FirstPicture   string `json:"firstPicture"`
	Description    string `json:"description"`
	Content        string `json:"content"`
	Cate           int    `json:"cate"`
	TagList        []int  `json:"tagList"`
	Words          string `json:"words"`
	ReadTime       int    `json:"readTime"`
	Views          int    `json:"views"`
	Appreciation   bool   `json:"appreciation"`
	Recommend      bool   `json:"recommend"`
	CommentEnabled bool   `json:"commentEnabled"`
	Top            bool   `json:"top"`
	Published      bool   `json:"published"`
	Password       string `json:"password"`
}

type CreateBlogRes struct {
}

type ShowBlogsReq struct {
	g.Meta        `path:"/admin/blogs" method:"get" summart:"展示文章列表" tags:"展示文章列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Title         string `json:"title"`
	CategoryId    int    `json:"categoryId"`
	PageNum       int    `json:"pageNum"`
	PageSize      int    `json:"pageSize"`
}

type ShowBlogsRes struct {
	Blogs      Blogs        `json:"blogs"`
	Categories []Categories `json:"categories"`
}

type BlogDetailReq struct {
	g.Meta        `path:"/admin/blog" method:"get" summart:"文章详情" tags:"文章详情"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
}

type BlogDetailRes struct {
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

type UpdateBlogTopReq struct {
	g.Meta        `path:"/admin/blog/top" method:"put" summart:"修改文章置顶信息" tags:"修改文章置顶信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
	Top           bool   `json:"top"`
}

type UpdateBlogTopRes struct {
}

type UpdateBlogRecommendReq struct {
	g.Meta        `path:"/admin/blog/recommend" method:"put" summart:"修改文章推荐信息" tags:"修改文章推荐信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
	Top           bool   `json:"top"`
}

type UpdateBlogRecommendRes struct {
}

type UpdateBlogVisibilityReq struct {
	g.Meta         `path:"/admin/blog/{Id}/visibility" method:"put" summart:"修改文章可见性信息" tags:"修改文章可见性信息"`
	Authorization  string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id             int    `json:"id" in:"path" dc:"文章id"`
	Appreciation   bool   `json:"appreciation" dc:"赞赏"`
	Recommend      bool   `json:"recommend" dc:"推荐"`
	CommentEnabled bool   `json:"commentEnabled" dc:"评论"`
	Top            bool   `json:"top" dc:"置顶"`
	Published      bool   `json:"published" dc:"是否公开"`
	Password       string `json:"password" dc:"私密密码"`
}

type UpdateBlogVisibilityRes struct {
}

type UpdateBlogReq struct {
	g.Meta        `path:"/admin/blog" method:"put" summart:"更新文章" tags:"更新文章"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Id            int    `json:"id"`
	CreateBlogReq
}

type UpdateBlogRes struct {
}

type DeleteBlogReq struct {
	g.Meta `path:"/admin/blog" method:"delete" summart:"删除文章" tags:"删除文章"`
	Id     int `json:"id"`
}

type DeleteBlogRes struct {
}

type NbBlogsReq struct {
	g.Meta  `path:"/blogs" method:"get" summart:"前台展示文章列表" tags:"前台展示文章列表"`
	PageNum int `json:"PageNum"`
}

type NbBlogsRes struct {
	List      []BlogBase `json:"list"`
	Totalpage int        `json:"totalPage"`
}
