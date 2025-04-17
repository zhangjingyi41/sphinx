import axios from 'axios'
import { ElMessage } from 'element-plus'

// 创建axios实例
const instance = axios.create({
    baseURL: '/api', // 使用/api前缀，由Vite代理转发到实际后端
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json'
    }
})

// 请求拦截器
instance.interceptors.request.use(
    config => {
        // 可以在这里添加token等认证信息
        const token = localStorage.getItem('token')
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

// 响应拦截器
instance.interceptors.response.use(
    response => {
        // 检查响应数据中的code字段
        const res = response.data
        if (res.code !== undefined && res.code !== 0) {
            // code不为0，视为业务逻辑错误
            ElMessage.error(res.msg || '请求失败')
            return Promise.reject(res)
        }
        // code为0或未定义，视为成功
        return response
    },
    error => {
        // HTTP错误处理
        if (error.response) {
            const res = error.response.data
            // 如果错误响应中包含msg字段，优先使用它
            if (res && res.msg) {
                ElMessage.error(res.msg)
            } else {
                // 否则根据HTTP状态码给出通用错误消息
                switch (error.response.status) {
                    case 401:
                        // 未授权，可以处理登出逻辑
                        ElMessage.error('未授权，请重新登录')
                        break
                    case 404:
                        ElMessage.error('请求的资源不存在')
                        break
                    case 500:
                        ElMessage.error('服务器错误')
                        break
                    default:
                        ElMessage.error(`请求错误: ${error.response.status}`)
                }
            }
        } else {
            ElMessage.error('网络错误，请检查您的网络连接')
        }
        return Promise.reject(error)
    }
)

export default instance 