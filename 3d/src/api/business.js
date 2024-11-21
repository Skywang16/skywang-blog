import {
	request
} from "@/utils/request.js";

/**
 * 示例
 * 名称:'URL' // 接口功能简述
 */
const API = {
	newsList: `article/news/list`, // 资讯列表
	website: `setting/website/detail`, // 网站配置
	protocol:`setting/protocol/detail`, // 政策协议
	keyWordsLists: `article/keyWords/list`, // 文章关键词列表
	cateList: `common/album/cateList`, // 相册
	albumList: `common/album/albumList` // 相册
};


export function cateList(pramter) {
	return request({
		url: API.cateList,
		params: pramter,
		method: "GET",
	});
}
export function albumList(pramter) {
	return request({
		url: API.albumList,
		params: pramter,
		method: "GET",
	});
}
// 资讯列表
export function newsList(pramter) {
	return request({
		url: API.newsList,
		params: pramter,
		method: "GET",
	});
}

// 武器列表
export function weaponList(pramter) {
	return request({
		url: API.weaponList,
		params: pramter,
		method: "GET",
	});
}

// 角色列表
export function heroList(pramter) {
	return request({
		url: API.heroList,
		params: pramter,
		method: "GET",
	});
}

// 武器列表
export function website(pramter) {
	return request({
		url: API.website,
		params: pramter,
		method: "GET",
	});
}

// 地图列表
export function mapList(pramter) {
	return request({
		url: API.mapList,
		params: pramter,
		method: "GET",
	});
}


// 政策协议
export function protocol(pramter) {
	return request({
		url: API.protocol,
		params: pramter,
		method: "GET",
	});
}

export function keyWordsLists(pramter) {
	return request({
		url: API.keyWordsLists,
		params: pramter,
		method: "GET",
	});
}




