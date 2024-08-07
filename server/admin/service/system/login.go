package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"runtime/debug"
	"strconv"
	"time"
)

type ISystemLoginService interface {
	Login(c *gin.Context, req *req.SystemLoginReq) (res resp.SystemLoginResp, e error)
	Logout(req *req.SystemLogoutReq) (e error)
	RecordLoginLog(c *gin.Context, adminId uint, username string, errStr string) (e error)
}

// NewSystemLoginService 初始化
func NewSystemLoginService(db *gorm.DB, adminSrv ISystemAuthAdminService) ISystemLoginService {
	return &systemLoginService{db: db, adminSrv: adminSrv}
}

// systemLoginService 系统登录服务实现类
type systemLoginService struct {
	db       *gorm.DB
	adminSrv ISystemAuthAdminService
}

// Login 登录
func (loginSrv systemLoginService) Login(c *gin.Context, req *req.SystemLoginReq) (res resp.SystemLoginResp, e error) {
	sysAdmin, err := loginSrv.adminSrv.FindByUsername(req.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	} else if err != nil {
		core.Logger.Errorf("Login FindByUsername err: err=[%+v]", err)
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.Failed.Msg()); e != nil {
			return
		}
		e = response.Failed
		return
	}
	if sysAdmin.IsDelete == 1 {
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	}
	if sysAdmin.IsDisable == 1 {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.LoginDisableError.Msg()); e != nil {
			return
		}
		e = response.LoginDisableError
		return
	}
	md5Pwd := util.ToolsUtil.MakeMd5(req.Password + sysAdmin.Salt)
	if sysAdmin.Password != md5Pwd {
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			// 自定义类型
			case response.RespType:
				panic(r)
			// 其他类型
			default:
				core.Logger.Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
				loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.Failed.Msg())
				panic(response.Failed)
			}
		}
	}()
	token := util.ToolsUtil.MakeToken()
	adminIdStr := strconv.FormatUint(uint64(sysAdmin.ID), 10)

	//非多处登录
	if sysAdmin.IsMultipoint == 0 {
		sysAdminSetKey := config.AdminConfig.BackstageTokenSet + adminIdStr
		ts := util.RedisUtil.SGet(sysAdminSetKey)
		if len(ts) > 0 {
			var keys []string
			for _, t := range ts {
				keys = append(keys, t)
			}
			util.RedisUtil.Del(keys...)
		}
		util.RedisUtil.Del(sysAdminSetKey)
		util.RedisUtil.SSet(sysAdminSetKey, token)
	}

	// 缓存登录信息
	util.RedisUtil.Set(config.AdminConfig.BackstageTokenKey+token, adminIdStr, 7200)
	loginSrv.adminSrv.CacheAdminUserByUid(sysAdmin.ID)
	// 更新登录信息
	err = loginSrv.db.Model(&sysAdmin).Updates(
		system.SystemAuthAdmin{LastLoginIp: c.ClientIP(), LastLoginTime: time.Now().Unix()}).Error
	if err != nil {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.SystemError.Msg()); e != nil {
			return
		}
		if e = response.CheckErr(err, "Login Updates err"); e != nil {
			return
		}
	}
	// 记录登录日志
	if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, ""); e != nil {
		return
	}
	// 返回登录信息
	return resp.SystemLoginResp{Token: token}, nil
}

// Logout 退出
func (loginSrv systemLoginService) Logout(req *req.SystemLogoutReq) (e error) {
	util.RedisUtil.Del(config.AdminConfig.BackstageTokenKey + req.Token)
	return
}

// RecordLoginLog 记录登录日志并更新访问量
func (loginSrv systemLoginService) RecordLoginLog(c *gin.Context, adminId uint, username string, errStr string) (e error) {
	ua := core.UAParser.Parse(c.GetHeader("user-agent"))
	var status uint8
	if errStr == "" {
		status = 1
	}
	// 获取当前日期
	today := time.Now().UTC().Format("2006-01-02")
	// 检查是否存在当前日期的记录
	var authLog system.SystemAuthLog
	result := loginSrv.db.Where("create_time = ?", today).First(&authLog)
	if result.Error != nil {
		// 如果记录不存在，遍历表中每一天并计算总访问量
		var totalVisits int64
		var allLogs []system.SystemAuthLog
		loginSrv.db.Find(&allLogs)
		for _, log := range allLogs {
			totalVisits += log.TodayVisits
		}
		// 创建新记录
		authLog = system.SystemAuthLog{
			CreateTime:  time.Now().UTC(),
			TodayVisits: 1,
			TotalVisits: totalVisits + 1,
			UpdateTime:  time.Now().UTC(),
		}
		err := loginSrv.db.Create(&authLog).Error
		if err != nil {
			return response.CheckErr(err, "RecordLoginLog Create err")
		}
	} else {
		// 如果记录存在，增加访问量
		authLog.TodayVisits++
		authLog.TotalVisits++
		authLog.UpdateTime = time.Now().UTC()
		err := loginSrv.db.Model(&authLog).Where("create_time = ?", today).Updates(&authLog).Error
		if err != nil {
			return response.CheckErr(err, "RecordLoginLog Save err")
		}
	}
	// 记录登录日志
	err := loginSrv.db.Create(&system.SystemLogLogin{
		AdminId:  adminId,
		Username: username,
		Ip:       c.ClientIP(),
		Os:       ua.Os.Family,
		Browser:  ua.UserAgent.Family,
		Status:   status,
	}).Error
	if err != nil {
		return response.CheckErr(err, "RecordLoginLog Create err")
	}
	return nil
}
