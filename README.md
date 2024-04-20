# 毕业设计

#### 介绍
无畏契约小程序，小程序后台及其服务

#### 软件架构
软件架构说明

#### 环境要求

服务端

运行环境	要求版本	推荐版本
Go	>=1.16	1.18
Mysql	>=5.7	5.7
nginx 或 apache	无限制	-
Redis	>=6.0	7.0.0

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
cd likeadmin_go/server
* 下载依赖
go mod tidy
go mod download

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

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
