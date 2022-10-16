package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// Get @Summary 获取一个标签
//@Produce json
//@Param id  path  int  true "标签id"
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {

}

// List @Summary 获取多个标签
//@Produce json
//@Param name body string false "标签名称" minlength(3) maxlength(100)
//@Param state query int false "状态" Enums(0,1) default 1
//@Param page query int false "页码"
//@Param page_size query int false "每页数量"
//@Success 200 {object} model.TagSwagger "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
    param:= struct {
		Name string `form:"name",binding:"max=100"`
		State string `form:"state,default=1",binding:"oneof=0 1"`
	}{}
    response:=app.NewResponse(c)
    valid,errs:=app.BindAndValid(c,&param)
    if !valid{
    	global.Logger.Errorf("App.BindAndValid errs:%v",errs)
    	errResp:=errcode.InvalidParams.WithDetails(errs.Errors()...)
    	response.ToErrorResponse(errResp)
		return
	}
	response.ToResponse(gin.H{})
}

// Update @Summary 更新标签
//@Produce json
//@Param id   path  int  true "标签id"
//@Param name body string false "标签名称" minlength(3) maxlength(100)
//@Param state body int  false "状态" Enums(0,1) default 1
//@Param modified_by body string true "修改者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {

}


// Create @Summary 新增标签
//@Produce json
//@Param name body string true "标签名称" minlength(3) maxlength(100)
//@Param state body int false "状态" Enums(0,1) default 1
//@Param created_by body string true "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}


// Delete @Summary 删除标签
//@Produce json
//@Param   id  path  int  true "标签id"
//@Success 200 {object} model.Tag "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

}
