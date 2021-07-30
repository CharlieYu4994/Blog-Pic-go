# Blog-Pic-go
![Go](https://img.shields.io/github/go-mod/go-version/CharlieYu4994/Blog-Pic-go)![GPLv3](https://img.shields.io/github/license/CharlieYu4994/Blog-Pic-go) ![GitHub last commit](https://img.shields.io/github/last-commit/CharlieYu4994/Blog-Pic-go)

![Demo](https://api.lassi-cat.cn:60443/bingapi/bing)

~~一个潦草的 Bing 每日一图 API 的 Go 语言实现，并提供图片地址持久化存储~~

一个潦草的博客壁纸的 Go 语言实现，并提供图片地址持久化存储

> PS：已经不只是 Bing 每日一图 API 了……


## 特点
+ 使用 Golang 编写，SQLite3 作为数据库实现，较为轻量化
+ 提供自定义设置项，方便个性化
+ 提供 Bing 美图以及 NASA 的 APOD


## 使用

### 环境要求
+ 任意一台服务器 (required)
+ Go 编译环境 (required)

### 编译
1. 在 [Release](https://github.com/CharlieYu4994/Blog-Pic-go/releases) 页面下载源码
2. 解压并进入文件夹
3. 在文件夹下运行 `go build ./`

### 部署
1. 根据『config.json.template』创建配置文件
2. 将数据库、配置文件、二进制文件放在同一目录下运行
3. 访问 `your.domain/bing` 查看效果


## 访问
访问 `your.domain/<type>?dat=<date>&res=<resolution>` 然后你会被重定向到图片

### 参数
1. type 图片类型 (required)
   + bing：Bing 美图
   + apod：NASA 的 APOD

1. dat 日期偏移 (optional)
   + 0：当天的，同理 1 为一天前的，最大为你设置的 `picnum`-1
   + -1：随机图片
   + 可以不提供，不提供默认为 `0`

2. res 分辨率 (optional)
   + hdres：1080p 分辨率的图片
   + uhdres：当天的最高分辨率，分辨率不定
   + `height`x`width`：自定义分辨率（**必须被 Bing 官方支持**）
   + 可以不提供，不提供默认为 `hdres`
   + NASA 的 APOD 不支持调整分辨率


## 资源

### 官方演示
大家可以随便使用
```
api.lassi-cat.cn:60443/bingapi/<type>?dat=<date>&res=<resolution>
```
更多详情请查看 [Demo 主页](https://bing.charlie.moe)
> 此 Demo 由 [LassiCat](https://github.com/LassiCat) 的服务器托管运行 ~~（其实就是咱的）~~

### 图像数据库
~~在『main』分支下的『picture.db』为从2021年1月1日开始的图像，不定期更新 (其实不更新)~~

请使用 [Blog-Pic-db](https://github.com/CharlieYu4994/Blog-Pic-db) 提供的数据

[点击](https://github.com/CharlieYu4994/Blog-Pic-db/releases/latest/download/picture.db) 以下载最新数据库

### 对 BingAPI 的研究
[对 BingAPI 的简单研究](https://blog.charlie.moe/post/BingAPI-00/)


## 版权
Bing 美图由微软 Bing 提供，APOD 由 NASA 提供，**使用请注意版权，不可用于商业用途**

本程序由 [@CharlieYu4994](https://blog.charlie.moe/) 编写，以 GPLv3 协议发布

**本程序不支持，不鼓励一切商业用途**


