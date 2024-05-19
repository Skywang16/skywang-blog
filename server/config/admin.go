package config

import "github.com/gin-gonic/gin"

// AdminConfig 后台公共配置
var AdminConfig = adminConfig{
	// 管理缓存键
	BackstageManageKey: "backstage:manage",
	// 角色缓存键
	BackstageRolesKey: "backstage:roles",
	// 令牌缓存键
	BackstageTokenKey: "backstage:token:",
	// 令牌的集合
	BackstageTokenSet: "backstage:token:set:",

	// 免登录验证
	NotLoginUri: []string{
		"system:login",           // 登录接口
		"common:index:config",    // 配置接口
		"common:upload:image",    // 文件上传
		"common:album:albumDel",  // 文件删除
		"system:admin:add",       // 人员新增
		"article:news:list",      // 文章列表
		"setting:website:detail", // 网站配置
	},

	// 免权限验证
	NotAuthUri: []string{
		"system:logout",     // 退出登录
		"system:menu:menus", // 系统菜单
		"system:menu:route", // 菜单路由
		"system:admin:self", // 管理员信息
	},

	// 演示模式白名单
	ShowWhitelistUri: []string{
		"system:login",  // 登录接口
		"system:logout", // 退出登录
	},

	// 请求临时数据
	SuperAdminId:   1,
	ReqAdminIdKey:  "admin_id",
	ReqRoleIdKey:   "role",
	ReqUsernameKey: "username",
	ReqNicknameKey: "nickname",
}

type adminConfig struct {
	BackstageManageKey string
	BackstageRolesKey  string
	BackstageTokenKey  string
	BackstageTokenSet  string
	NotLoginUri        []string
	NotAuthUri         []string
	ShowWhitelistUri   []string
	SuperAdminId       uint
	ReqAdminIdKey      string
	ReqRoleIdKey       string
	ReqUsernameKey     string
	ReqNicknameKey     string
}

func (cnf adminConfig) GetAdminId(c *gin.Context) uint {
	adminId, ok := c.Get(cnf.ReqAdminIdKey)
	if !ok {
		return 0
	}
	return adminId.(uint)
}

func (cnf adminConfig) GetRoleId(c *gin.Context) string {
	roleId, ok := c.Get(cnf.ReqRoleIdKey)
	if !ok {
		return ""
	}
	return roleId.(string)
}

func (cnf adminConfig) GetUsername(c *gin.Context) string {
	username, ok := c.Get(cnf.ReqUsernameKey)
	if !ok {
		return ""
	}
	return username.(string)
}

func (cnf adminConfig) GetNickname(c *gin.Context) string {
	nickname, ok := c.Get(cnf.ReqNicknameKey)
	if !ok {
		return ""
	}
	return nickname.(string)
}
