import request from "@/utils/request";

// 用户列表
export function getUserList(params: any) {
    return request.get({url: "/system/admin/list", params});
}

// 用户详情
export function getUserDetail(params: any) {
  return request.get({ url: "/user/detail", params });
}

/* // 用户编辑
export function userEdit(params: any) {
  return request.post({ url: "/system/admin/detail", params });
} */
