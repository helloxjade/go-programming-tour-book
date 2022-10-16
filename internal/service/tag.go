package service

//form tag标签代表表单的映射字段名 binding代表入参校验的规则名称
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"required,min=3,max=100"`
	State string `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     string `form:"state,default=1" binding:"oneof=0 1"`
}
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	ModefiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	State      string `form:"state" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
