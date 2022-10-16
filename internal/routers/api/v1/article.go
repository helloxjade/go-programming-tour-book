package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Article struct {

}

func NewArticle() Article{
	return Article{}
}

// Get @Summary 获取一篇文章
//@Produce json
//@Param  id  path  int  true "文章id"
//@Success 200 {object} model.Article "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router /api/v1/articles/{id} [get]
func (a Article)Get(c *gin.Context){
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return

}

// List @Summary 获取多篇文章
//@Produce json
//@Success 200 {object} model.ArticlesSwagger "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/articles [get]
func (a Article)List(c *gin.Context){

}

// Update @Summary 更新文章
//@Produce json
//@Param id   path  int  true "文章id"
//@Param state query int true "文章状态" Enums(0,1) default 1
//@Param modified_by body string true "更新者"minlength(3) maxlength(100)
//@Param title body string false "文章标题" minlength(3) maxlength(100)
//@Param desc body string false "文章描述"
//@Param content body string false "文章内容"
//@Param cover_image_url body string false "封面图片" minlength(3) maxlength(100)
//@Success 200 {object} model.Article "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /articles/{id} [put]
//@Router /articles/:id/state [patch]
func (a Article)Update(c *gin.Context){

}
// Create @Summary 新增文章
//@Produce json
//@Param state query int false "文章状态" Enums(0,1) default 1
//@Param created_by body string true "创建者" minlength(3) maxlength(100)
//@Param title body string true "文章标题" minlength(3) maxlength(100)
//@Param desc body string false "文章描述"
//@Param content body string true "文章内容"
//@Param cover_image_url body string false "封面图片" minlength(3) maxlength(100)
//@Success 200 {object} model.Article "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/tags [post]
func (a Article)Create(c *gin.Context){

}

// Delete @Summary 删除一篇文章
//@Produce json
//@Param  id  path  int  true "文章id"
//@Success 200 {object} model.Article "成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部失败"
//@Router  /api/v1/articles/{id} [delete]
func (a Article)Delete(c *gin.Context){

}
