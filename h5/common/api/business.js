import {
	request
} from "@/util/request.js";

/**
 * 示例
 * 名称:'URL' // 接口功能简述
 */
const API = {
	newsList: `article/news/list`, // 资讯列表
	weaponList: `information/weapons/list`, // 资讯列表
	website: `setting/website/detail`, // 网站配置
};

// 资讯列表
export function newsList(pramter) {
	return request({
		url: API.newsList,
		data: pramter,
		method: "GET",
	});
}

// 武器列表
export function weaponList(pramter) {
	return request({
		url: API.weaponList,
		data: pramter,
		method: "GET",
	});
}

// 武器列表
export function website(pramter) {
	return request({
		url: API.website,
		data: pramter,
		method: "GET",
	});
}


