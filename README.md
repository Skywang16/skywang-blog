# 个人博客

#### 介绍
个人博客、博客管理系统、博客小程序

#### 软件架构
软件架构说明

#### 环境要求

服务端

运行环境	要求版本	推荐版本
Go	>=1.16	1.18
Mysql	>=5.7	5.7
nginx 或 apache	无限制	-
Redis	>=6.0	7.0.0

PC端

运行环境	要求版本	推荐版本
Node.js	>=14.18.1	14.18.1

后台管理

运行环境	要求版本	推荐版本
Node.js	>=14.18.1	14.18.1

小程序

运行环境	要求版本	推荐版本
Node.js	>=14.18.1	14.18.1

#### 安装教程

服务端

* 安装依赖
项目推荐使用Go1.18版本，最低使用Go1.16
由于国内没法安装 go.org/x 包下面的东西，推荐使用 goproxy.cn 或者 goproxy.io
如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
* 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on
* 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
* 下载依赖
go mod tidy
go mod download

PC端
* npm install
* npm run dev

后台管理
* npm install
* npm run dev

小程序
* 使用Hbuilder X运行
* 使用Hbuilder菜单栏点击运行 == 运行到浏览器

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx


