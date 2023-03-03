# Piwriw Blog

<a href="#">
<img src="https://img.shields.io/badge/Vue-2x-green.svg">
</a>
<a href="#">
<img src="https://img.shields.io/badge/Gin--green.svg">
</a>
<a href="#">
<img src="https://img.shields.io/badge/Sqlx--green.svg">
</a>
<!-- PROJECT LOGO -->
<br />

<p align="center">

  <a href="https://github.compiwriw/piwirw_blog/">
    <img src="https://s1.ax1x.com/2023/03/02/ppFLbUx.png" alt="Logo" width="550" height="350">
  </a>

<h3 align="center">基于Gin的个人博客系统</h3>
  <p align="center">
    一个后端使用了Gin、数据库ORM使用Sqlx，前端使用了Vue2.X的个人博客系统
    <br />
    <a href="https://github.com/Piwriw/piwriw_blog/blob/master/README.md"><strong>探索本项目的文档 »</strong></a>
    <br />

[comment]: <> (    <br />)

[comment]: <> (    <a href="https://github.compiwriw/piwirw_blog">查看Demo</a>)

[comment]: <> (    ·)

[comment]: <> (    <a href="https://github.compiwriw/piwirw_blog/issues">报告Bug</a>)

[comment]: <> (    ·)

[comment]: <> (    <a href="https://github.compiwriw/piwirw_blog/issues">提出新特性</a>)
  </p>


本篇README.md面向开发者

## 目录

- [项目概述](#项目概述)
- [页面演示](#页面演示)
    - [用户端](#用户端)
    - [管理端](#管理端)
- [快速启动](#快速启动)
- [文件目录说明](#后端)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 项目概述

Piwriw_blog是一个博客系统，前端基于Vue2构建了用户端和管理员系统，后端基于Gin和Sqlx构建。

### 页面演示

#### 用户端

- 首页

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302200947.png" />
</div>

- 详情

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230303102800.png" />
</div>

#### 管理端

- 文章编辑

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302220716.png" />
</div>

- 文章列表

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302220812.png" />
</div>

- 评论管理

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302222146.png" />
</div>

- 分类列表

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302222105.png" />
</div>

- 用户列表

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302222631.png" />
</div>

- 个人设置

<div align="center">
  <img  width="92%" style="border-radius:10px;margin-top:20px;margin-bottom:20px;box-shadow: 2px 0 6px gray;" src="https://cdn.jsdelivr.net/gh/Piwriw/PicGo/image/20230302222607.png" />
</div>

### 快速启动
- 1. 克隆项目到本地
```sh
git clone git@github.com:Piwriw/piwriw_blog.git
```
- 2. 下载后端依赖
     (**前提**：你已经开启了go modules 并且设置了国内镜像代理)
```shell
cd piwriw_blog 
go mod tidy
```
- 3. 下载前端依赖
    
```shell
cd piwriw_blog_web/front
yarn install
cd piwriw_blog_web/admin
yarn install
```

- 4. 启动项目
    Go version=1.18
```shell
go run main.go
yarn run serve
```

### 文件目录说明
#### 后端
``` 
│  go.mod 
│  go.sum
│  main.go  //gin入口文件
│  web_app.log
│      
├─conf
│      dev.yaml //配置文件，数据库，JWT等等
│      
├─controller
│      admin.go 
│      article.go
│      category.go
│      code.go
│      comment.go
│      profile.go
│      request.go
│      response.go
│      user.go
│      validator.go //validator库的配置
│      
├─dao
│  └─mysql      
│          admin.go
│          article.go
│          category.go
│          comment.go
│          error_code.go
│          mysql.go
│          profile.go
│          user.go
│          
├─logger
│      logger.go    //logger的配置
│      
├─middlewares
│      auth.go      // jwt中间件认证
│      cors.go      // 跨域的解决
│      
├─models
│  │  article.go
│  │  category.go
│  │  comments.go
│  │  params.go
│  │  Po.go
│  │  profile.go
│  │  time.go
│  │  user.go
│  │  
│  └─response
│          article.go
│          category.go
│          comment.go
│          page.go
│          user.go
│          
├─pkg
│  └─jwt
│          jwt.go
│          
├─router
│      admin.go
│      article.go
│      category.go
│      comment.go
│      profile.go
│      route.go
│      user.go
│      
├─service
│      admin.go
│      article.go
│      category.go
│      comment.go
│      profile.go
│      user.go
│      
├─setting
│      setting.go   //配置项读取
│      
├─sql
│      my_blog.sql
│      my_blog_withdata.sql


        


```

### 使用到的框架

- [Gin](https://gin-gonic.com/zh-cn/)
- [Vue](https://cn.vuejs.org/)


### 作者

[Piwriw](https://github.com/Piwriw)

@Email:piwriw@163.com

前端使用了 wejectchan的[ginblog](https://gitee.com/wejectchan/ginblog?_from=gitee_search) 并作出修改

### 鸣谢

- [wejectchan](https://gitee.com/wejectchan)


<!-- links -->



