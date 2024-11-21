import {
	request
} from "@/utils/request.js";

/**
 * 示例
 * 名称:'URL' // 接口功能简述
 */
const API = {
	token: `system/login`, 
	register: `system/admin/add`,
	delImg:`common/album/albumDel`, //删除图片
	getUserInfo: `system/admin/self`, // 用户信息
};

export function token(pramter) {
	return request({
		url: API.token,
		params: pramter,
		method: "POST",
	});
}

export function register(pramter) {
	return request({
		url: API.register,
		params: pramter,
		method: "POST",
	});
}

export function delImg(pramter) {
	return request({
		url: API.delImg,
		params: pramter,
		method: "POST",
	});
}

export function getUserInfo(pramter) {
	return request({
		url: API.getUserInfo,
		params: pramter,
		method: "GET",
	});
}