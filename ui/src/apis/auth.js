import request from './request'

// export const Login = (params = {}) => request.post('/login', { params })

export const Login = payload => request.post('/login/', payload)
