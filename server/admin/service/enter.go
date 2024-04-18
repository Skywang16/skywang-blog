package service

import (
	"likeadmin/admin/service/article"
	"likeadmin/admin/service/common"
	"likeadmin/admin/service/information"
	"likeadmin/admin/service/setting"
	"likeadmin/admin/service/system"
)

var InitFunctions = []interface{}{
	// common
	common.NewAlbumService,
	common.NewIndexService,
	common.NewUploadService,
	// setting
	setting.NewSettingCopyrightService,
	setting.NewSettingDictDataService,
	setting.NewSettingDictTypeService,
	setting.NewSettingProtocolService,
	setting.NewSettingStorageService,
	setting.NewSettingWebsiteService,
	// system
	system.NewSystemAuthAdminService,
	system.NewSystemAuthDeptService,
	system.NewSystemAuthMenuService,
	system.NewSystemAuthPermService,
	system.NewSystemAuthRoleService,
	system.NewSystemLoginService,
	system.NewSystemLogsServer,
	// article
	article.NewArticleNewsService,
	// information
	information.NewInformationWeaponsService,
	information.NewInformationWeaponsSeriesService,
	information.NewInformationHeroService,
	information.NewInformationHeroTypeService,
}
