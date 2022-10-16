# **日志库**

gopkg.in/natefinch/lumberjack.v2

**核心功能：**
把日志写进滚动文件中，允许我们设置日志文件的最大占用空间、最大生存周期、可保留的最多旧文件数，如果有出现超出设置项的情况，则对日志进行滚动处理。
可以减免一些文件操作类的代码编写，把核心逻辑摆在日志标准化处理上。

_Swagger_
==
Swagger 工具集可根据OpenAPI规范生成各类与接口相关联的工具

_常见的流程：_
---

编写注解 -》 调用生成库 》 生``成标准描述文件 -》生成/导入对应的Swagger工具

    `go get -u github.com/swaggo/swag/cmd/swag@v1.6.5`
    `go get -u github.com/swaggo/gin-swagger@v1.2.0`
    `go get -u github.com/swaggo/files`
    `go get -u github.com/alecthomas/template`
    验证安装成功 
    `swag -v`
    生成swagger文档
    `swag init`
    会在docs目录下生成 swagger.json及swagger.yaml及docs.go文件

    源码中的doc.json相当于一个内部标识，可以读取生成的swagger注解

    访问swagger文档
    只需在routers中进行默认初始化和注册对应的路由，见代码
    主要是初始化docs包和注册一个针对swagger的路由
    浏览器访问：http://127.0.0.1:8000/swagger/index.html
    包含项目主体信息、接口路由、模型信息


Swagger注解
--

    @Summary 摘要
    @Produce API可以产生的MIME类型的列表。可以理解为相应类型，如json,XML,HTML
    @Success 响应成功 从左到右分别为 ：状态码、参数类型、数据类型和注释
    @Param 参数格式 分别为: 状态码、入参类型、数据类型、是否必填和注释
    @Failure 响应失败，分别为：状态码、参数类型、数据类型和注释
    @Router 路由 路由地址和HTTP方法

接口校验validator
------
    gin框架的内部的模型绑定和验证默认使用的就是go-playground/validator
    标签与含义：
    len: 长度要求与len给定的一致
国际化
==
    简单的国际化可以通过中间件配合语言包的方式来实现
- github.com/go-playground/locales：多语言包，从CLDR项目生成的一组多语言环境，主要在il8n软件包中使用，需要与universal-translator配套使用
- go-playground/universal-translator ：通用翻译器，是一个使用CLDR数据+复数规则的GO语言il8n转换器
- go-playground/validator/v10/translations：validator翻译器