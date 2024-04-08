'use strict';
const db = uniCloud.database();
exports.main = async (event, context) => {
	res = await db.collection('uni-cms').get()
	//返回数据给客户端
	return res
};
