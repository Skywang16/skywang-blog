import request from "@/utils/request";

// 文章列表
export function articleLists(params?: any) {
    return request.get({url: "/article/news/list", params});
}

// 文章列表
export function articleAll(params?: any) {
    return request.get({url: "/article/news/all", params});
}

// 添加文章
export function articleAdd(params: any) {
    return request.post({url: "/article/news/add", params});
}

// 编辑文章
export function articleEdit(params: any) {
    return request.post({url: "/article/news/edit", params});
}

// 删除文章
export function articleDelete(params: any) {
    return request.post({url: "/article/news/del", params});
}

// 文章详情
export function articleDetail(params: any) {
    return request.get({url: "/article/news/detail", params});
}

// 文章关键词列表
export function keyWordsLists(params?: any) {
    return request.get({ url: "/article/keyWords/list", params });
  }
  
  // 添加文章关键词
  export function keyWordsAdd(params: any) {
    return request.post({ url: "/article/keyWords/add", params });
  }
  
  // 编辑文章关键词
  export function keyWordsEdit(params: any) {
    return request.post({ url: "/article/keyWords/edit", params });
  }
  
  // 删除文章关键词
  export function keyWordsDelete(params: any) {
    return request.post({ url: "/article/keyWords/del", params });
  }
  
  // 文章关键词详情
  export function keyWordsDetail(params: any) {
    return request.get({ url: "/article/keyWords/detail", params });
  }
  