# Bing-API-go
一个潦草的 Bing 每日一图 API 的 Go 语言实现，并提供图片地址持久化存储

使用 GPL3.0 开源协议

## 特点
+ 使用 Golang 编写，SQLite3 作为数据库实现，较为轻量化
+ 提供自定义设置项，方便个性化

## 使用
访问 `your.domain/bing?dat=<date>&res=<resolution>` 然后你会被重定向到图片

### 参数
1. dat (optional)
   + 0：一天前的，同理 1 为一天前的，最大为你设置的 `picnum`-1
   + -1：随机图片
   + 可以不提供，不提供默认为 `0`

2. res (optional)
   + hdres：1080p 分辨率的图片
   + uhdres：当天的最高分辨率，分辨率不定
   + <height>x<width>：自定义分辨率（必须被 Bing 官方支持）
   + 可以不提供，不提供默认为 `hdres`


## 环境要求
+ 任意一台服务器 (required)
+ Go 编译环境 (optional 仅在需要自行编译时需要)

## 部署

### 有你的平台的二进制文件
1. 在 『[Relrase](https://github.com/CharlieYu4994/Bing-API-go/releases)』页面下载二进制文件
2. 根据『config.json.template』创建配置文件
3. 将数据库、配置文件、二进制文件放在同一目录下运行
4. 访问 `your.domain/bing` 查看效果

### 没有你的平台的二进制文件
1. 在 『[Relrase](https://github.com/CharlieYu4994/Bing-API-go/releases)』页面下载源码
2. 解压并进入文件夹
3. 在文件夹下运行 `go build ./`
4. 根据『config.json.template』创建配置文件
5. 将数据库、配置文件、二进制文件放在同一目录下运行
6. 访问 `your.domain/bing` 查看效果

## 资源

### 官方演示
大家可以随便使用
```
bing-api.lassi-cat.cn:60443/bing
```
![Demo](https://bing-api.lassi-cat.cn:60443/bing)

### 对 BingAPI 的研究
https://blog.charlie.moe/post/BingAPI-00/



