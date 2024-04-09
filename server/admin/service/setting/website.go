package setting

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/core/response"
	"likeadmin/util"
	"strings"
)

type ISettingWebsiteService interface {
	Detail() (res map[string]string, e error)
	Save(wsReq req.SettingWebsiteReq) (e error)
}

// NewSettingWebsiteService 初始化
func NewSettingWebsiteService(db *gorm.DB) ISettingWebsiteService {
	return &settingWebsiteService{db: db}
}

// settingWebsiteService 网站信息配置服务实现类
type settingWebsiteService struct {
	db *gorm.DB
}

// Detail 获取网站信息
func (wSrv settingWebsiteService) Detail() (res map[string]string, e error) {
	data, err := util.ConfigUtil.Get(wSrv.db, "website")
	if e = response.CheckErr(err, "Detail Get err"); e != nil {
		return
	}
	homeBannerItems := strings.Split(data["homeBanner"], ",")
	var absoluteHomeBannerUrls []string
	for _, item := range homeBannerItems {
		absoluteUrl := util.UrlUtil.ToAbsoluteUrl(strings.TrimSpace(item))
		absoluteHomeBannerUrls = append(absoluteHomeBannerUrls, absoluteUrl)
	}
	homeBanner := strings.Join(absoluteHomeBannerUrls, ",")
	res = map[string]string{
		"name":       data["name"],
		"logo":       util.UrlUtil.ToAbsoluteUrl(data["logo"]),
		"favicon":    util.UrlUtil.ToAbsoluteUrl(data["favicon"]),
		"backdrop":   util.UrlUtil.ToAbsoluteUrl(data["backdrop"]),
		"homeBanner": homeBanner,
	}
	return res, nil
}

// Save 保存网站信息
func (wSrv settingWebsiteService) Save(wsReq req.SettingWebsiteReq) (e error) {
	var urls []string
	// Process homeBanner URLs
	bannerItems := strings.Split(wsReq.HomeBanner, ",")
	for _, item := range bannerItems {
		parsedUrl := util.UrlUtil.ToAbsoluteUrl(strings.TrimSpace(item))
		urls = append(urls, parsedUrl)
	}
	wsReq.HomeBanner = strings.Join(urls, ",")
	// Save website info
	err := util.ConfigUtil.Set(wSrv.db, "website", "name", wsReq.Name)
	if e = response.CheckErr(err, "Save Set name err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "logo", util.UrlUtil.ToRelativeUrl(wsReq.Logo))
	if e = response.CheckErr(err, "Save Set logo err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "favicon", util.UrlUtil.ToRelativeUrl(wsReq.Favicon))
	if e = response.CheckErr(err, "Save Set favicon err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "backdrop", util.UrlUtil.ToRelativeUrl(wsReq.Backdrop))
	if e = response.CheckErr(err, "Save Set backdrop err"); e != nil {
		return
	}
	err = util.ConfigUtil.Set(wSrv.db, "website", "homeBanner", wsReq.HomeBanner)
	if e = response.CheckErr(err, "Save Set homeBanner err"); e != nil {
		return
	}
	return
}
