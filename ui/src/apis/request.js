import axios from 'axios'
import { BACK_API, TOKEN_NAME } from '@/config'
import router from '@/router'
import storge from '@/utils/storage'
import Vue from 'vue'

const vm = new Vue()
const request = axios.create({
  baseURL: BACK_API,
  timeout: 3 * 60 * 1000, // 3 minutes
  responseType: 'json',
})

request.interceptors.request.use((config) => {
  const token = storge.getItem(TOKEN_NAME)
  Object.defineProperty(config.headers, 'Authorization', {
    value: `Bearer ${token}`,
    configurable: true,
    enumerable: true,
    writable: true,
  })
  return config
})

request.interceptors.response.use(
  (response) => response.data,
  (error) => {
    if (error.response && error.response.status === 401) {
      router.push({ name: 'Login' })
      return new Promise(() => ({})) // pending的promise，中止promise链
    } else if (error.response && error.response.status === 403) {
      vm.$Message.error('您没有权限进行此操作')
    } else if (error.response && error.response.status === 500) {
      vm.$Message.error('服务器或网络错误')
    } else if (error.code === 'ECONNABORTED') {
      vm.$Message.error('网络连接超时')
    } else if (error.response && error.response.status === 400) {
      const msg = typeof error.response.data === 'string' // 某些接口的报错直接为字符串，比如上传
        ? error.response.data
        : Object.values(error.response.data).join('；')
      vm.$Message.error(msg)
    } else {
      vm.$Message.error(Object.values(error.response.data).join('；'))
    }

    return Promise.reject(error.response)
  },
)

export default request
