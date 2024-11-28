import axios from 'axios'

 let BaseUrl = 'https://api.skywang.asia/api/'
// let BaseUrl = 'http://127.0.0.1:8001/api/'

// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: BaseUrl,
  timeout: 30000 // 请求超时时间
})

// 异常拦截处理器
const errorHandler = (error) => {
  if (error.response) {
    const { data = {}, status } = error.response
    // 403 无权限
    if (status === 403) {
      console.log('您没有相关访问权限')
    }
    // 401 未登录/未授权
    if (status === 401 && !(data.result && data.result.isLogin)) {
      console.log('认证失败')
    }
    // 自定义错误处理
    if (data.code === 333) {
      console.log('登录超时，请重新登录')
    }
  }
  return Promise.reject(error)
}

// 请求拦截器
const requestHandler = (config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers['version'] = '1.3.3'
  }
  config.headers['Content-Type'] = config.params ? config.params.contentType : 'application/json'
  return config
}

// Add a request interceptor
request.interceptors.request.use(requestHandler, errorHandler)

// 响应拦截器
const responseHandler = (response) => {
  return response.data
}

// Add a response interceptor
request.interceptors.response.use(responseHandler, errorHandler)

// 导出
export { request, BaseUrl }
