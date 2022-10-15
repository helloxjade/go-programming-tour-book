package model



//创建公共model
type Model struct {
	ID         uint32 `gorm:"primary_key"   json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
//创建标签
type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag)TableName() string {
	return "blog_tag"
}

//创建文章
type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

func (a Article)TableName() string {
	return "blog_article"
}


//创建文章标签关联表
type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}


func (a ArticleTag)TableName() string {
	return "blog_article_tag"
}