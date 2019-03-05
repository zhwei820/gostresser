import request from './request'

export const fetchHealthCheck = (params = {}) => request.get('status', { params })
