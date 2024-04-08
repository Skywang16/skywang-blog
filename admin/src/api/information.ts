import request from "@/utils/request";

// 武器列表
export function weaponsLists(params?: any) {
    return request.get({url: "/information/weapons/list", params});
}

// 添加武器
export function weaponsAdd(params: any) {
    return request.post({url: "/information/weapons/add", params});
}

// 编辑武器
export function weaponsEdit(params: any) {
    return request.post({url: "/information/weapons/edit", params});
}

// 删除武器
export function weaponsDelete(params: any) {
    return request.post({url: "/information/weapons/del", params});
}

// 武器详情
export function weaponsDetail(params: any) {
    return request.get({url: "/information/weapons/detail", params});
}

// 武器系列列表
export function weaponsSeriesLists(params?: any) {
    return request.get({url: "/information/weaponsSeries/list", params});
}

// 添加武器系列
export function weaponsSeriesAdd(params: any) {
    return request.post({url: "/information/weaponsSeries/add", params});
}

// 编辑武器系列
export function weaponsSeriesEdit(params: any) {
    return request.post({url: "/information/weaponsSeries/edit", params});
}

// 删除武器系列
export function weaponsSeriesDelete(params: any) {
    return request.post({url: "/information/weaponsSeries/del", params});
}

// 武器系列详情
export function weaponsSeriesDetail(params: any) {
    return request.get({url: "/information/weaponsSeries/detail", params});
}