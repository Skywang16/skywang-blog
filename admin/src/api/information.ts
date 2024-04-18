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

// 角色列表
export function heroLists(params?: any) {
    return request.get({url: "/information/hero/list", params});
}

// 添加角色
export function heroAdd(params: any) {
    return request.post({url: "/information/hero/add", params});
}

// 编辑角色
export function heroEdit(params: any) {
    return request.post({url: "/information/hero/edit", params});
}

// 删除角色
export function heroDelete(params: any) {
    return request.post({url: "/information/hero/del", params});
}

// 角色详情
export function heroDetail(params: any) {
    return request.get({url: "/information/hero/detail", params});
}

// 角色定位列表
export function heroTypeLists(params?: any) {
    return request.get({url: "/information/heroType/list", params});
}

// 添加角色定位
export function heroTypeAdd(params: any) {
    return request.post({url: "/information/heroType/add", params});
}

// 编辑角色定位
export function heroTypeEdit(params: any) {
    return request.post({url: "/information/heroType/edit", params});
}

// 删除角色定位
export function heroTypeDelete(params: any) {
    return request.post({url: "/information/heroType/del", params});
}

// 角色定位详情
export function heroTypeDetail(params: any) {
    return request.get({url: "/information/heroType/detail", params});
}