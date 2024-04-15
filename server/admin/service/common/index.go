package common

import (
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/model/system"
	"likeadmin/util"
	"time"
)

type IIndexService interface {
	Console() (res map[string]interface{}, e error)
	Config() (res map[string]interface{}, e error)
}

// NewIndexService 初始化
func NewIndexService(db *gorm.DB) IIndexService {
	return &indexService{db: db}
}

// indexService 主页服务实现类
type indexService struct {
	db *gorm.DB
}

// Console 控制台数据
func (iSrv indexService) Console() (res map[string]interface{}, e error) {
	// 版本信息
	name, err := util.ConfigUtil.GetVal(iSrv.db, "website", "name", "LikeAdmin-Go")
	if e = response.CheckErr(err, "Console Get err"); e != nil {
		return
	}
	version := map[string]interface{}{
		"name":    name,
		"version": config.Config.Version,
		"website": "www.likeadmin.cn",
		"based":   "Vue3.x、ElementUI、MySQL",
		"channel": map[string]string{
			"gitee": "https://gitee.com/ciaovita/graduation-design",
		},
	}
	// 从数据库中获取访问量数据
	var authLog system.SystemAuthLog
	// 这里假设您有一个方法来查询最新的 SystemAuthLog 记录
	err = iSrv.db.Order("create_time desc").First(&authLog).Error
	if e = response.CheckErr(err, "Failed to get SystemAuthLog"); e != nil {
		return
	}
	var totalUsers int64
	if err := iSrv.db.Model(&system.SystemAuthAdmin{}).Where("role = ?", "2").Count(&totalUsers).Error; err != nil {
		// 处理错误
	}
	// 今日数据
	today := map[string]interface{}{
		"time":        time.Now().Format("2006-01-02 15:04"), // 使用当前时间，精确到分钟
		"todayVisits": authLog.TodayVisits,                   // 从数据库中获取当天访问量
		"totalVisits": authLog.TotalVisits,                   // 从数据库中获取总访问量
		"todayUsers":  authLog.TodayUsers,                    // 这里假设您没有记录新增用户，可能需要添加逻辑来获取
		"totalUsers":  totalUsers,                            // 从数据库中获取总用户数
	}
	// 访客图表
	now := time.Now()
	var date []string
	var visits []int
	for i := 14; i >= 0; i-- {
		date = append(date, now.AddDate(0, 0, -i).Format(core.DateFormat))
		// 查询数据库获取当天的访问量数据
		var authLog system.SystemAuthLog
		err := iSrv.db.Where("create_time >= ? AND create_time < ?", now.AddDate(0, 0, -i).Format("2006-01-02 00:00:00"), now.AddDate(0, 0, -i+1).Format("2006-01-02 00:00:00")).First(&authLog).Error
		if err != nil {
			// 处理错误
		}
		// 将int64转换为int
		visits = append(visits, int(authLog.TodayVisits))
	}
	visitor := map[string]interface{}{
		"date": date,
		"list": visits,
	}
	return map[string]interface{}{
		"version": version,
		"today":   today,
		"visitor": visitor,
	}, nil
}

// Config 公共配置
func (iSrv indexService) Config() (res map[string]interface{}, e error) {
	website, err := util.ConfigUtil.Get(iSrv.db, "website")
	if e = response.CheckErr(err, "Config Get err"); e != nil {
		return
	}
	copyrightStr, err := util.ConfigUtil.GetVal(iSrv.db, "website", "copyright", "")
	if e = response.CheckErr(err, "Config GetVal err"); e != nil {
		return
	}
	var copyright []map[string]string
	if copyrightStr != "" {
		err = util.ToolsUtil.JsonToObj(copyrightStr, &copyright)
		if e = response.CheckErr(err, "Config JsonToObj err"); e != nil {
			return
		}
	} else {
		copyright = []map[string]string{}
	}
	return map[string]interface{}{
		"webName":     website["name"],
		"webLogo":     util.UrlUtil.ToAbsoluteUrl(website["logo"]),
		"webFavicon":  util.UrlUtil.ToAbsoluteUrl(website["favicon"]),
		"webBackdrop": util.UrlUtil.ToAbsoluteUrl(website["backdrop"]),
		"ossDomain":   config.Config.PublicUrl,
		"copyright":   copyright,
	}, nil
}
