export function request(params) {
	const BASE_URL = "http://localhost:8001/api/";
	uni.setStorageSync("BASE_URL", BASE_URL);
	let BASE_HEADER = {
		"content-type": params.contentType || "application/json",
	};
	// 获取token
	let token = uni.getStorageSync("token");
	if (token) {
		BASE_HEADER.Token = token;
		BASE_HEADER.version = '1.3.3'
	}
	let url = params.url;
	let method = params.method;
	let header = params.header || BASE_HEADER;
	let data = params.data || {};
	let dataType = params.dataType || "json"
	// 请求方式 GET / POST
	if (method) {
		// 转大写
		method = method.toUpperCase();
	}
	// 发起网络请求
	return new Promise((resolve, reject) => {
		uni.request({
			url: BASE_URL + url,
			method: method || "POST",
			header: header,
			data: data,
			dataType: dataType,
			timeout: 60000,
			success: (res) => {
				resolve(res); // 返回成功状态
			},
			fail: (err) => {
				reject(err); // 返回失败状态
				uni.hideLoading();
				typeof params.fail == "function" && params.fail(err.data);
			},
		});
	});
}