# 智多汀——基于ChatGLM的AI答题应用平台

![Bitbucket open issues](https://img.shields.io/bitbucket/issues/iuricode/README-template?style=for-the-badge)
![Bitbucket open pull requests](https://img.shields.io/bitbucket/pr-raw/iuricode/README-template?style=for-the-badge)

[//]: # (<img src="imagem.png" alt="Exemplo imagem">)

> 该项目属于大学生对go语言的后端开发的练习项目，项目的目的是为了实现一个基于ChatGLM的AI答题应用平台，该平台可以实现用户注册、登录、答题、查看答题记录等功能。

### 当前任务进度

该AI答题应用平台是一款根据大模型的推断、文本概括以及扩写等
AIGC（AI-Generated-Content）能力结合UGC（User-Generated-Content）规
则的生成式应用使用及分享平台。用户可以基于AI快速制作并发布多种答题应用、
支持检索、在线答题并基于评分算法或AI得到问答总结，并进行统计分析:

- [x] 用户登录、注册、修改个人信息、查看个人信息、管理员权限等
- [ ] App应用管理、发布、编辑、删除、查看、搜索等
- [ ] AI智能出题、智能测评
- [ ] 用户发布App、分享、评论、点赞、收藏等
- [ ] 用户答题、查看答题记录、查看答题结果、查看答题统计等

## 💻 技术栈

作为一名go语言的后端开发者，该项目使用了以下技术栈:

- Gin: 一个用Go（Golang）编写的Web框架。它具有类似于Martini的API，但性能更好.
- Gorm: 一个用Go（Golang）编写的ORM库，旨在对数据库进行简单且快速的操作.
- Viper: 一个用Go（Golang）编写的配置库，支持多种配置文件格式，如JSON、TOML、YAML、HCL、envfile和Java属性文件.
- JWT-go: 一个用Go（Golang）编写的JWT库，用于生成和验证JWT.
- Go-redis: 一个用Go（Golang）编写的Redis客户端库，支持多种命令.
- Go-MySQL-Driver: 一个用Go（Golang）编写的MySQL驱动库，用于连接MySQL数据库.

## 🚀 安装

由于go语言的跨平台特性，该项目可以在Linux、macOS和Windows上运行。在安装之前，你需要安装go语言的开发环境，具体安装方法可以参考[go语言官网](https://golang.org/doc/install)。

```
# git clone 项目
git clone xxxx://xxxx/xxxx.git
```
```
cd goAiroject
# 安装依赖
go get -d ./...
```
```
# 运行项目
go build
./goAiroject
```


## 📝 License

该项目已获得许可。有关更多详细信息，请参阅[LICENSE]（LICENSE.md）文件。
