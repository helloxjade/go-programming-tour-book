package service

type CountArticlesRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticlesListRequest struct {
	Name  string `form:"name" binding:"required,min=3,max=100"`
	State string `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateArticlesRequest struct {
	Name          string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         string `form:"state,default=1" binding:"oneof=0 1"`
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc"`
	Content       string `form:"content" binding:"required"`
	CoverImageUrl string `form:"cover_image_url"`
}
type UpdateArticlesRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	ModefiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	State      string `form:"state" binding:"oneof=0 1"`
}

type DeleteArticlesRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
